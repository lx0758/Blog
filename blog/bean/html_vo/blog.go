package html_vo

type BlogVO struct {
	SiteDomain         string
	SiteTitle          string
	SiteDescription    string
	SiteKeywords       string
	SiteBeianIcp       string
	SiteBeianPs        string
	SiteBaidu          string
	SiteCreateYear     string
	SiteArticleCount   int
	SiteCategoryCount  int
	SiteTagCount       int
	OwnerAvatar        string
	OwnerNickname      string
	OwnerDescription   string
	OwnerEmail         string
	OwnerGithub        string
	OwnerWeibo         string
	OwnerGoogle        string
	OwnerTwitter       string
	OwnerFacebook      string
	OwnerStackOverflow string
	OwnerYoutube       string
	OwnerInstagram     string
	OwnerSkype         string
	Links              []LinkVO
}
