package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"pibbo/config"
	"pibbo/utils"
	"sort"
	"strings"
	"time"
)

type Time time.Time

type Article struct {
	Title       string   `json:"title"`
	Date        Time     `json:"date"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Author      string   `json:"author"`
	MusicId     string   `json:"musicId"`
	Path        string
	ShortUrl    string
	Category    string
	TotalViews  int
}

type Articles []Article

type ArticleDetail struct {
	Article
	Body string
}

func initArticlesAndImages(dir string) (Articles, map[string]string, error) {
	var articles Articles
	shortUrlMap := make(map[string]string)

	articles, err := RecursiveReadArticles(dir)
	if err != nil {
		return articles, shortUrlMap, err
	}
	sort.Sort(articles)
	for i := len(articles) - 1; i >= 0; i-- {
		//这里必须使用倒序的方式生成 shortUrl,因为如果有相同的文章标题，
		// 倒序会将最老的文章优先生成shortUrl，保证和之前的 shortUrl一样
		article := articles[i]
		keyword := utils.Generator(utils.CHARSET_RANDOM_ALPHANUMERIC, article.Title, func(url, keyword string) bool {
			//保证 keyword 唯一
			_, ok := shortUrlMap[keyword]
			return !ok
		})
		articles[i].ShortUrl = keyword
		shortUrlMap[keyword] = article.Path
	}
	return articles, shortUrlMap, nil
}

func ArticleSearch(articles *Articles, search string, category string) Articles {

	var articleList Articles
	for _, article := range *articles {

		pass := true

		if search != "" && strings.Index(article.Title, search) == -1 {
			pass = false
		}
		if category != "" && strings.Index(article.Category, category) == -1 {
			pass = false
		}
		if pass {
			articleList = append(articleList, article)
		}

	}
	return articleList
}

func RecursiveReadArticles(dir string) (Articles, error) {

	var articles Articles

	dirInfo, err := os.Stat(dir)

	if err != nil {
		return articles, err
	}
	if !dirInfo.IsDir() {
		return articles, errors.New("目标不是一个目录")
	}

	fileOrDir, err := ioutil.ReadDir(dir)
	if err != nil {
		return articles, err
	}

	for _, fileInfo := range fileOrDir {
		name := fileInfo.Name()
		path := dir + "/" + name
		upperName := strings.ToUpper(name)
		if fileInfo.IsDir() {
			subArticles, err := RecursiveReadArticles(path)
			if err != nil {
				return articles, err
			}
			articles = append(articles, subArticles...)
			// 判断是否.MD结尾
		} else if strings.HasSuffix(upperName, ".MD") {
			article, err := ReadArticle(path)
			if err != nil {
				return articles, err
			}
			articles = append(articles, article)
			// 判断是否为图片
		} else if strings.HasSuffix(upperName, ".PNG") ||
			strings.HasSuffix(upperName, ".GIF") ||
			strings.HasSuffix(upperName, ".JPG") {

			dst := config.Cfg.CurrentDir + "/images/" + name
			log.Println(utils.IsFile(dst))
			if !utils.IsFile(dst) {
				// 拷贝图片到目的路径
				_, _ = utils.CopyFile(path, dst)
			}
		}

	}
	return articles, nil
}

func ReadArticle(path string) (Article, error) {
	article, _, err := readMarkdown(path)
	if err != nil {
		return article, err
	}
	return article, nil
}

func ReadArticleDetail(path string) (ArticleDetail, error) {
	_, articleDetail, err := readMarkdown(path)
	if err != nil {
		return articleDetail, err
	}
	articleDetail.Article.TotalViews += 1
	return articleDetail, nil
}

func readMarkdown(path string) (Article, ArticleDetail, error) {
	var article Article
	var articleDetail ArticleDetail
	mdFile, err := os.Stat(path)

	if err != nil {
		return article, articleDetail, err
	}
	if mdFile.IsDir() {
		return article, articleDetail, errors.New("this path is Dir")
	}
	markdown, err := ioutil.ReadFile(path)

	if err != nil {
		return article, articleDetail, err
	}
	// 去掉所有空格
	markdown = bytes.TrimSpace(markdown)

	article.Path = path
	// 文章分类
	article.Category = GetCategoryName(path)
	// 文章标题
	article.Title = strings.TrimSuffix(strings.ToUpper(mdFile.Name()), ".MD")
	// 文章时间
	article.Date = Time(mdFile.ModTime())

	// 不是以```json开头的情况下判断
	if !bytes.HasPrefix(markdown, []byte("```json")) {
		article.Description = cropDesc(markdown)
		articleDetail.Article = article
		articleDetail.Body = string(markdown)
		return article, articleDetail, nil
	}

	// 把```json 替换成空
	markdown = bytes.Replace(markdown, []byte("```json"), []byte(""), 1)
	// 以```为分隔符切成2个子切片
	markdownArrInfo := bytes.SplitN(markdown, []byte("```"), 2)
	article.Description = cropDesc(markdownArrInfo[1])

	if err := json.Unmarshal(bytes.TrimSpace(markdownArrInfo[0]), &article); err != nil {
		article.Title = "文章[" + article.Title + "]解析 JSON 出错，请检查。"
		article.Description = err.Error()
		return article, articleDetail, nil
	}
	article.Path = path
	article.Title = strings.ToUpper(article.Title)

	articleDetail.Article = article
	articleDetail.Body = string(markdownArrInfo[1])
	return article, articleDetail, nil
}

func cropDesc(c []byte) string {
	content := []rune(string(c))
	contentLen := len(content)

	if contentLen <= config.Cfg.DescriptionLen {
		return string(content[0:contentLen])
	}

	return string(content[0:config.Cfg.DescriptionLen])
}

func (t *Time) UnmarshalJSON(b []byte) error {
	date, err := time.ParseInLocation(`"`+config.Cfg.TimeLayout+`"`, string(b), time.Local)
	if err != nil {
		return nil
	}
	*t = Time(date)
	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {

	return []byte(t.Format(`"` + config.Cfg.TimeLayout + `"`)), nil
}

func (t Time) Format(layout string) string {
	return time.Time(t).Format(layout)
}

func (a Articles) Len() int { return len(a) }

func (a Articles) Less(i, j int) bool { return time.Time(a[i].Date).After(time.Time(a[j].Date)) }

func (a Articles) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
