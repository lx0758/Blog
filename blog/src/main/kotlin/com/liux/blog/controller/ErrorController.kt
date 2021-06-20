package com.liux.blog.controller

import com.liux.blog.bean.Resp
import com.liux.blog.config.ExceptionConfig
import com.liux.blog.service.ThemeService
import com.liux.blog.toBean
import com.liux.blog.toJSONString
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.autoconfigure.web.ServerProperties
import org.springframework.boot.autoconfigure.web.servlet.error.BasicErrorController
import org.springframework.boot.web.error.ErrorAttributeOptions
import org.springframework.boot.web.servlet.error.ErrorAttributes
import org.springframework.http.MediaType
import org.springframework.http.ResponseEntity
import org.springframework.stereotype.Controller
import org.springframework.ui.Model
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.servlet.ModelAndView
import javax.servlet.http.HttpServletRequest
import javax.servlet.http.HttpServletResponse

@Controller
@RequestMapping("\${server.error.path:\${error.path:/error}}")
class ErrorController(
    errorAttributes: ErrorAttributes,
    serverProperties: ServerProperties,
): BasicErrorController(errorAttributes, serverProperties.error) {

    @Autowired
    private lateinit var themeService: ThemeService

    override fun getErrorPath(): String {
        return errorProperties.path
    }

    /**
     * 处理网页请求的错误
     */
    override fun errorHtml(request: HttpServletRequest, response: HttpServletResponse): ModelAndView {
        return super.errorHtml(request, response).apply {
            val template = when (response.status) {
                404 -> "404"
                else -> "error"
            }
            viewName = themeService.render(ErrorModel(this), template)
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
        return ResponseEntity(bodyMap, ExceptionConfig.transformHttpStatus(httpStatus))
    }

    class ErrorModel(
            private val modelAndView: ModelAndView
    ) : Model {

        override fun addAttribute(attributeName: String, attributeValue: Any?): Model {
            modelAndView.addObject(attributeName, attributeValue)
            return this
        }

        override fun addAttribute(attributeValue: Any): Model {
            modelAndView.addObject(attributeValue)
            return this
        }

        override fun addAllAttributes(attributeValues: MutableCollection<*>): Model {
            attributeValues.forEach {
                if (it != null) modelAndView.addObject(it)
            }
            return this
        }

        override fun addAllAttributes(attributes: MutableMap<String, *>): Model {
            modelAndView.addAllObjects(attributes)
            return this
        }

        override fun mergeAttributes(attributes: MutableMap<String, *>): Model {
            modelAndView.addAllObjects(attributes)
            return this
        }

        override fun containsAttribute(attributeName: String): Boolean {
            return modelAndView.model.containsKey(attributeName)
        }

        override fun getAttribute(attributeName: String): Any? {
            return modelAndView.model[attributeName]
        }

        override fun asMap(): MutableMap<String, Any> {
            return modelAndView.model
        }

    }
}
