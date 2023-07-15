package com.liux.blog.bean.vo.api

class DashboardVO(
    val articleCount: Int,
    val categoryCount: Int,
    val tagCount: Int,
    val uploadCount: Int,
    val commentCount: Int,
    val browseCount: Int,

    val newArticles: List<ArticleVO>,
    val newComments: List<CommentVO>,
    val newUploads: List<FileVO>,
)