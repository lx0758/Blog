package com.liux.blog.config

import jakarta.servlet.RequestDispatcher
import jakarta.servlet.http.HttpServletRequest
import jakarta.servlet.http.HttpServletResponse
import org.apache.commons.logging.LogFactory
import org.apache.tomcat.util.http.fileupload.impl.SizeLimitExceededException
import org.springframework.http.HttpStatus
import org.springframework.web.bind.annotation.*
import org.springframework.web.client.HttpClientErrorException
import org.springframework.web.multipart.MaxUploadSizeExceededException
import org.springframework.web.servlet.resource.NoResourceFoundException

@ControllerAdvice
class ExceptionConfig {

    private val logger = LogFactory.getLog(javaClass)

    @ExceptionHandler(value = [Exception::class])
    fun handlerException(
        request: HttpServletRequest,
        response: HttpServletResponse,
        exception: Exception
    ): String {
        logger.error("${exception.javaClass.simpleName}: ${exception.message} (${request.requestURI})")

        val httpStatus = when(exception) {
            is SizeLimitExceededException, is MaxUploadSizeExceededException -> HttpStatus.PAYLOAD_TOO_LARGE
            is HttpClientErrorException -> HttpStatus.valueOf(exception.statusCode.value())
            is NoResourceFoundException -> HttpStatus.NOT_FOUND
            else -> HttpStatus.INTERNAL_SERVER_ERROR
        }

        request.setAttribute(RequestDispatcher.ERROR_EXCEPTION, exception)
        request.setAttribute(RequestDispatcher.ERROR_EXCEPTION_TYPE, exception::class.java)
        request.setAttribute(RequestDispatcher.ERROR_MESSAGE, exception.message ?: httpStatus.reasonPhrase)
        request.setAttribute(RequestDispatcher.ERROR_REQUEST_URI, request.requestURI)
        request.setAttribute(RequestDispatcher.ERROR_STATUS_CODE, httpStatus.value())

        return "forward:/error"
    }
}