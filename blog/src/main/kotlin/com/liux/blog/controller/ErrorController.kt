package com.liux.blog.controller

import com.liux.blog.service.ThemeService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.autoconfigure.web.ServerProperties
import org.springframework.boot.autoconfigure.web.servlet.error.BasicErrorController
import org.springframework.boot.web.error.ErrorAttributeOptions
import org.springframework.boot.web.servlet.error.ErrorAttributes
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

    override fun error(request: HttpServletRequest): ResponseEntity<MutableMap<String, Any>> {
        return super.error(request).apply {

        }
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
            return modelAndView.model.get(attributeName)
        }

        override fun asMap(): MutableMap<String, Any> {
            return modelAndView.model
        }

    }
}
