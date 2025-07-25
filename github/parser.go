package github

import (
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Parse(r io.Reader) ([]Repo, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	var repos []Repo

	doc.Find("main article").Each(func(_ int, s *goquery.Selection) {
		// Find owner and name
		href := s.Find("h2 a").AttrOr("href", "")
		parts := strings.Split(href, "/")
		if len(parts) != 3 { //nolint: mnd // href format: /owner/name
			return
		}

		// Find description
		description := s.Children().Get(2).FirstChild.Data //nolint: mnd // description is the 3rd child
		description = strings.TrimSpace(description)

		// Find language color
		languageColor := s.Find("span.repo-language-color").AttrOr("style", "")
		languageColor = strings.TrimPrefix(languageColor, "background-color:")
		languageColor = strings.TrimSpace(languageColor)

		// Find language
		language := s.Find(".d-inline-block span[itemprop='programmingLanguage']").Text()

		// Find stars
		starsTotal := s.Find(".Link--muted.d-inline-block.mr-3").First().Text()
		starsTotal = strings.TrimSpace(starsTotal)

		// Find forks
		forks := s.Find(".Link--muted.d-inline-block.mr-3").Eq(1).Text()
		forks = strings.TrimSpace(forks)

		// Find stars since
		starsSince := s.Find(".d-inline-block.float-sm-right").Text()
		starsSince = strings.TrimSpace(starsSince)

		repos = append(repos, Repo{
			Owner:         parts[1],
			Name:          parts[2],
			Description:   description,
			Language:      language,
			LanguageColor: languageColor,
			StarsTotal:    starsTotal,
			Forks:         forks,
			StarsSince:    starsSince,
		})
	})

	return repos, nil
}
