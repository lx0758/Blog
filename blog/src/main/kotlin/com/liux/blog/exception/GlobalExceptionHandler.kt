package com.liux.blog.exception

import org.springframework.web.bind.annotation.ControllerAdvice
import org.springframework.web.bind.annotation.ExceptionHandler
import javax.servlet.http.HttpServletResponse

@ControllerAdvice
class GlobalExceptionHandler {

    @ExceptionHandler(value = [HttpNotFoundException::class])
    fun handlerNotFound(response: HttpServletResponse, exception: HttpNotFoundException) {
        response.sendError(exception.rawStatusCode, exception.message)
    }
}
