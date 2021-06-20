package com.liux.blog.config

import com.liux.blog.bean.Resp
import com.liux.blog.toFormatJSONString
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import org.springframework.http.HttpMethod
import org.springframework.http.MediaType
import springfox.documentation.builders.*
import springfox.documentation.oas.annotations.EnableOpenApi
import springfox.documentation.schema.ScalarType
import springfox.documentation.service.ApiDescription
import springfox.documentation.service.ParameterType
import springfox.documentation.spi.DocumentationType
import springfox.documentation.spi.service.ApiListingScannerPlugin
import springfox.documentation.spi.service.contexts.DocumentationContext
import springfox.documentation.spring.web.plugins.Docket
import springfox.documentation.spring.web.readers.operation.CachingOperationNameGenerator

@EnableOpenApi
@Configuration
class SwaggerConfig : ApiListingScannerPlugin {

    @Bean
    fun createRestApi(): Docket {
        return Docket(DocumentationType.OAS_30)
            .apiInfo(
                ApiInfoBuilder()
                    .title("接口文档")
                    .description("博客后台接口文档")
                    .version("1.0")
                    .build()
            )
            .select()
            .apis(RequestHandlerSelectors.any())
            .paths(
                PathSelectors.ant("/api/**")
                    .or(PathSelectors.ant("/article/**/comment"))
            )
            .build()
    }

    override fun supports(delimiter: DocumentationType): Boolean {
        return true
    }

    override fun apply(context: DocumentationContext): MutableList<ApiDescription> {
        val responses = arrayListOf(
            ResponseBuilder()
                .code("200")
                .description("OK")
                .examples(arrayListOf(
                    ExampleBuilder().id("example-1").mediaType(MediaType.APPLICATION_JSON_VALUE).value(Resp.succeed().toFormatJSONString()).build(),
                ))
                .build(),
            ResponseBuilder().code("401").description("Unauthorized").build(),
            ResponseBuilder().code("403").description("Forbidden").build(),
            ResponseBuilder().code("404").description("Not Found").build(),
        )
        val sessionApiDescription = ApiDescriptionBuilder(context.operationOrdering())
            .path("/api/session")
            .operations(
                arrayListOf(
                    OperationBuilder(CachingOperationNameGenerator())
                        .method(HttpMethod.POST)
                        .summary("login")
                        .uniqueId("login")
                        .tags(setOf("ApiFilter"))
                        .requestParameters(
                            arrayListOf(
                                RequestParameterBuilder()
                                    .query {
                                        it.model { it.maybeConvertToScalar(ScalarType.STRING) }
                                    }
                                    .`in`(ParameterType.QUERY)
                                    .parameterIndex(0)
                                    .name("username")
                                    .description("username")
                                    .required(true)
                                    .build(),
                                RequestParameterBuilder()
                                    .query {
                                        it.model { it.maybeConvertToScalar(ScalarType.STRING) }
                                    }
                                    .`in`(ParameterType.QUERY)
                                    .parameterIndex(1)
                                    .name("password")
                                    .description("password")
                                    .required(true)
                                    .build(),
                            )
                        )
                        .responses(responses)
                        .build(),
                    OperationBuilder(CachingOperationNameGenerator())
                        .method(HttpMethod.DELETE)
                        .summary("logout")
                        .uniqueId("logout")
                        .tags(setOf("ApiFilter"))
                        .produces(setOf(MediaType.APPLICATION_JSON_VALUE))
                        .responses(responses)
                        .build()
                )
            )
            .build()

        return arrayListOf(
            sessionApiDescription,
        )
    }
}