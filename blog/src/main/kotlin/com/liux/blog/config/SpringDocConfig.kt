package com.liux.blog.config

import io.swagger.v3.oas.models.OpenAPI
import io.swagger.v3.oas.models.Operation
import io.swagger.v3.oas.models.PathItem
import io.swagger.v3.oas.models.info.Info
import io.swagger.v3.oas.models.media.Content
import io.swagger.v3.oas.models.media.MediaType
import io.swagger.v3.oas.models.media.Schema
import io.swagger.v3.oas.models.parameters.Parameter
import io.swagger.v3.oas.models.responses.ApiResponse
import io.swagger.v3.oas.models.responses.ApiResponses
import org.springdoc.core.GroupedOpenApi
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration

@Configuration
class SpringDocConfig {

    @Bean
    fun blogOpenApi(): OpenAPI {
        return OpenAPI()
            .info(
                Info()
                    .title("Blog API")
                    .description("Blog API Docment")
            )
    }

    @Bean
    fun blogApi(): GroupedOpenApi {
        return GroupedOpenApi.builder()
            .group("blog")
            .pathsToMatch("/article/**")
            .build()
    }

    @Bean
    fun adminApi(): GroupedOpenApi {
        val respApiResponses = ApiResponses()
            .addApiResponse(
                "200",
                ApiResponse()
                    .description("OK")
                    .content(Content().addMediaType("*/*", MediaType().schema(Schema<Any>().`$ref`("#/components/schemas/RespObject"))))
            )
        return GroupedOpenApi.builder()
            .group("admin")
            .pathsToMatch("/api/**")
            .addOpenApiCustomiser {
                it
                    .path(
                        "/api/session",
                        PathItem()
                            .post(
                                Operation()
                                    .tags(arrayListOf("api-login-filter"))
                                    .operationId("login")
                                    .addParametersItem(
                                        Parameter()
                                            .name("username")
                                            .`in`("query")
                                            .required(true)
                                            .schema(Schema<String>().apply {
                                                type = "string"
                                            })
                                    )
                                    .addParametersItem(
                                        Parameter()
                                            .name("password")
                                            .`in`("query")
                                            .required(true)
                                            .schema(Schema<String>().apply {
                                                type = "string"
                                            })
                                    )
                                    .addParametersItem(
                                        Parameter()
                                            .name("verifyCode")
                                            .`in`("query")
                                            .required(false)
                                            .schema(Schema<String>().apply {
                                                type = "string"
                                            })
                                    )
                                    .responses(respApiResponses)
                            )
                            .delete(
                                Operation()
                                    .tags(arrayListOf("api-login-filter"))
                                    .operationId("logout")
                                    .responses(respApiResponses)
                            )
                    )
            }
            .build()
    }
}