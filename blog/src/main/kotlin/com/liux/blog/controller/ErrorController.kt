package com.liux.blog.controller

import com.liux.blog.bean.Resp
import com.liux.blog.service.ThemeService
import com.liux.blog.toBean
import com.liux.blog.toJSONString
import jakarta.servlet.http.HttpServletRequest
import jakarta.servlet.http.HttpServletResponse
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.autoconfigure.web.ServerProperties
import org.springframework.boot.autoconfigure.web.servlet.error.BasicErrorController
import org.springframework.boot.web.error.ErrorAttributeOptions
import org.springframework.boot.web.servlet.error.ErrorAttributes
import org.springframework.http.HttpStatus
import org.springframework.http.MediaType
import org.springframework.http.ResponseEntity
import org.springframework.stereotype.Controller
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.servlet.ModelAndView

@Controller
@RequestMapping("\${server.error.path:\${error.path:/error}}")
class ErrorController(
    errorAttributes: ErrorAttributes,
    serverProperties: ServerProperties,
): BasicErrorController(errorAttributes, serverProperties.error) {

    @Autowired
    private lateinit var themeService: ThemeService

    /**
     * 处理网页请求的错误
     */
    override fun errorHtml(request: HttpServletRequest, response: HttpServletResponse): ModelAndView {
        return super.errorHtml(request, response).apply {
            val template = when (response.status) {
                404 -> "404"
                else -> "error"
            }
            viewName = themeService.render(template, model)
            addObject("error", getErrorAttributes(request, ErrorAttributeOptions.defaults()))
        }
    }

    /**
     * 处理接口请求的错误
     */
    override fun error(request: HttpServletRequest): ResponseEntity<MutableMap<String, Any>> {
        val httpStatus = getStatus(request)
        val errorAttributes = getErrorAttributes(request, getErrorAttributeOptions(request, MediaType.ALL))
        val body = Resp.failed(httpStatus.value(), (errorAttributes["message"] ?: httpStatus.reasonPhrase) as String)
        val bodyMap = body.toJSONString().toBean<MutableMap<String, Any>>()
        return ResponseEntity(bodyMap, HttpStatus.OK)
    }
}
