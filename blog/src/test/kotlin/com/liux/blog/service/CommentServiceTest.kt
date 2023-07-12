package com.liux.blog.service

import org.junit.jupiter.api.Test
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.test.context.SpringBootTest
import org.springframework.jdbc.core.ArgumentPreparedStatementSetter
import org.springframework.jdbc.core.JdbcTemplate

@SpringBootTest
class CommentServiceTest {
    @Autowired
    private lateinit var commentService: CommentService
    @Autowired
    private lateinit var locationService: LocationService
    @Autowired
    private lateinit var jdbcTemplate: JdbcTemplate

    @Test
    fun generateLocationInfo() {
        val comments = commentService.listByAdmin(
            articleId = null,
            author = null,
            email = null,
            ip = null,
            content = null,
            status = null,
            pageNum = 1,
            pageSize = Int.MAX_VALUE,
            orderName = null,
            orderMethod = null,
        )
        comments.forEach { comment ->
            val location = locationService.getLocationFromIp(comment.ip!!)
            jdbcTemplate.update(
                "UPDATE t_comment SET location=? WHERE id=?",
                ArgumentPreparedStatementSetter(arrayOf(location, comment.id))
            )
        }
    }
}