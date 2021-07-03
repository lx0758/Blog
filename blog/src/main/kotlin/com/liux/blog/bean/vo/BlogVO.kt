package com.liux.blog.bean.vo

import com.liux.blog.bean.po.*

data class BlogVO(
    var siteDomain: String,
    var siteTitle: String,
    var siteDescription: String,
    var siteKeywords: String,

    var siteBeian: String,
    var siteBaidu: String,
    var siteCreateYear: String,

    var siteArticleCount: Int,
    var siteCategoryCount: Int,
    var siteTagCount: Int,

    var ownerAvatar: String,
    var ownerNickname: String,
    var ownerDescription: String,
    var ownerGithub: String,
    var ownerEmail: String,
    var ownerWeibo: String,
    var ownerGoogle: String,
    var ownerTwitter: String,
    var ownerFacebook: String,
    var ownerStackOverflow: String,
    var ownerYoutube: String,
    var ownerInstagram: String,
    var ownerSkype: String,

    var links: List<LinkVO> = emptyList(),
) {
    companion object {
        fun of(configs: Map<String, String?>, articleCount: Int, categoryCount: Int, tagCount: Int, user: User, links: List<Link>): BlogVO {
            return BlogVO(
                configs[Config.KEY_SITE_DOMAIN] ?: "domain.org",
                configs[Config.KEY_SITE_TITLE] ?: "Blog",
                configs[Config.KEY_SITE_DESCRIPTION] ?: "",
                configs[Config.KEY_SITE_KEYWORDS] ?: "",
                configs[Config.KEY_SITE_BEIAN] ?: "",
                configs[Config.KEY_SITE_BAIDU] ?: "",
                configs[Config.KEY_SITE_CREATE_YEAR] ?: "2020",

                articleCount,
                categoryCount,
                tagCount,

                "/blog/images/avatar.gif",
                user.nickname ?: "无名氏",
                user.description ?: "这个人很懒，什么也没留下。",
                user.github ?: "",
                user.email ?: "",
                "",
                "",
                "",
                "",
                "",
                "",
                "",
                "",

                links.map { LinkVO.of(it) }
            )
        }
    }
}
