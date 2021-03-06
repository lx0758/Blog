package com.liux.blog.bean.vo

import com.liux.blog.bean.po.*
import java.util.*

data class BaseVO(
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

    var masterAvatar: String,
    var masterNickname: String,
    var masterDescription: String,

    var masterGithub: String,
    var masterEmail: String,
    var masterWeibo: String,
    var masterGoogle: String,
    var masterTwitter: String,
    var masterFacebook: String,
    var masterStackOverflow: String,
    var masterYoutube: String,
    var masterInstagram: String,
    var masterSkype: String,

    var links: List<LinkVO> = emptyList(),
) {
    companion object {
        fun of(configs: Map<String, String?>, articleCount: Int, categoryCount: Int, tagCount: Int, user: User, links: List<Link>): BaseVO {
            return BaseVO(
                configs[SITE_DOMAIN] ?: "domain.org",
                configs[SITE_TITLE] ?: "Blog",
                configs[SITE_DESCRIPTION] ?: "",
                configs[SITE_KEYWORDS] ?: "",
                configs[SITE_BEIAN] ?: "",
                configs[SITE_BAIDU] ?: "",
                configs[SITE_CREATE_YEAR] ?: "2020",

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
