package com.liux.blog.config

import com.fasterxml.classmate.TypeResolver
import com.google.common.base.Predicates
import org.springframework.beans.factory.annotation.Value
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import org.springframework.http.HttpMethod
import org.springframework.http.MediaType
import springfox.documentation.builders.*
import springfox.documentation.schema.ModelRef
import springfox.documentation.service.ApiDescription
import springfox.documentation.service.Parameter
import springfox.documentation.spi.DocumentationType
import springfox.documentation.spi.service.ApiListingScannerPlugin
import springfox.documentation.spi.service.contexts.DocumentationContext
import springfox.documentation.spring.web.plugins.Docket
import springfox.documentation.spring.web.readers.operation.CachingOperationNameGenerator
import springfox.documentation.swagger2.annotations.EnableSwagger2

@Configuration
@EnableSwagger2
class SwaggerConfig : ApiListingScannerPlugin {

    @Value("\${springfox.documentation.enabled:false}")
    private var enabled: Boolean = false

    @Bean
    fun createRestApi(): Docket {
        return Docket(DocumentationType.SWAGGER_2)
            .enable(enabled)
            .apiInfo(
                ApiInfoBuilder()
                    .title("接口文档")
                    .description("博客后台接口文档")
                    .version("1.0")
                    .build()
            )
            .select()
            .apis(RequestHandlerSelectors.any())
            .paths(Predicates.or(
                PathSelectors.ant("/api/**"),
                PathSelectors.ant("/article/**/comment")
            ))
            .build()
    }

    override fun supports(delimiter: DocumentationType): Boolean {
        return true
    }

    override fun apply(context: DocumentationContext): MutableList<ApiDescription> {
        val responseMessages = setOf(
            ResponseMessageBuilder()
                .code(200)
                .message("OK")
                .responseModel(ModelRef("Resp«Void»"))
                .build(),
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
                        .parameters(arrayListOf(
                            createStringParameter(1, "username", true),
                            createStringParameter(2, "password", true),
                            createStringParameter(3, "verifyCode", false),
                        ))
                        .responseMessages(responseMessages)
                        .build(),
                    OperationBuilder(CachingOperationNameGenerator())
                        .method(HttpMethod.DELETE)
                        .summary("logout")
                        .uniqueId("logout")
                        .tags(setOf("ApiFilter"))
                        .produces(setOf(MediaType.APPLICATION_JSON_VALUE))
                        .responseMessages(responseMessages)
                        .build()
                )
            )
            .build()

        return arrayListOf(
            sessionApiDescription,
        )
    }

    private fun createStringParameter(order: Int, name: String, required: Boolean): Parameter {
        return ParameterBuilder()
            .order(order)
            .name(name)
            .parameterType("query")
            .description(name)
            .type(TypeResolver().arrayType(String::class.java))
            .parameterAccess("access")
            .modelRef(ModelRef("string"))
            .required(required)
            .build()
    }
}