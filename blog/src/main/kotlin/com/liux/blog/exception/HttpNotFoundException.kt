package com.liux.blog.exception

import org.springframework.http.HttpStatus
import org.springframework.web.client.HttpStatusCodeException

class HttpNotFoundException : HttpStatusCodeException(HttpStatus.NOT_FOUND, HttpStatus.NOT_FOUND.reasonPhrase)
