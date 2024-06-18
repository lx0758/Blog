import org.jetbrains.kotlin.gradle.tasks.KotlinCompile

plugins {
    id("org.springframework.boot") version "3.3.0"
    id("io.spring.dependency-management") version "1.1.5"
    kotlin("jvm") version "2.0.0"
    kotlin("plugin.spring") version "2.0.0"
    id("org.jetbrains.kotlin.plugin.noarg") version "2.0.0"
}

group = "com.liux.blog"
version = "0.0.1-SNAPSHOT"

java {
    sourceCompatibility = JavaVersion.VERSION_17
}

dependencies {
    implementation("org.jetbrains.kotlin:kotlin-reflect")
    implementation("org.jetbrains.kotlin:kotlin-stdlib-jdk8")

    implementation("org.springframework.boot:spring-boot-starter")
    implementation("org.springframework.boot:spring-boot-starter-web")
    implementation("org.springframework.boot:spring-boot-starter-mail")
    implementation("org.springframework.boot:spring-boot-starter-security")
    implementation("org.springframework.boot:spring-boot-starter-thymeleaf")

    implementation("org.jsoup:jsoup:1.17.2")
    implementation("org.lionsoul:ip2region:2.7.0")
    implementation("org.postgresql:postgresql:42.7.3")
    implementation("com.github.ua-parser:uap-java:1.6.1")
    implementation("com.fasterxml.jackson.module:jackson-module-kotlin:2.17.1")
    implementation("org.mybatis.spring.boot:mybatis-spring-boot-starter:3.0.3")
    implementation("com.github.pagehelper:pagehelper-spring-boot-starter:2.1.0")
    implementation("com.github.xiaoymin:knife4j-openapi3-jakarta-spring-boot-starter:4.5.0")

    implementation("org.commonmark:commonmark:0.22.0")
    implementation("org.commonmark:commonmark-ext-ins:0.22.0")
    implementation("org.commonmark:commonmark-ext-autolink:0.22.0")
    implementation("org.commonmark:commonmark-ext-heading-anchor:0.22.0")
    implementation("org.commonmark:commonmark-ext-task-list-items:0.22.0")
    implementation("org.commonmark:commonmark-ext-image-attributes:0.22.0")
    implementation("org.commonmark:commonmark-ext-gfm-tables:0.22.0")
    implementation("org.commonmark:commonmark-ext-gfm-strikethrough:0.22.0")

    developmentOnly("org.springframework.boot:spring-boot-devtools")
    developmentOnly("org.springframework.boot:spring-boot-starter-actuator")

    testImplementation("org.springframework.boot:spring-boot-starter-test")
    testImplementation("org.springframework.security:spring-security-test")
}

tasks.withType<KotlinCompile> {
    kotlinOptions {
        freeCompilerArgs = listOf("-Xjsr305=strict")
        jvmTarget = "17"
    }
}

tasks.withType<Test> {
    useJUnitPlatform()
}

// http://www.kotlincn.net/docs/reference/compiler-plugins.html
noArg {
    annotation("com.liux.blog.bean.NoArgsConstructor")
}

abstract class BuildAdminTask : AbstractExecTask<BuildAdminTask>(BuildAdminTask::class.java) {
    init {
        workingDir = File("../admin")
        commandLine = listOf("node", "./node_modules/vite/bin/vite.js", "build")
    }
}
tasks.register<BuildAdminTask>("buildAdmin").configure {
    group = "build"
}

tasks.register<Copy>("buildAdminToBlog").configure {
    group = "build"
    dependsOn("buildAdmin")

    from("../admin/dist")
    into("./src/main/resources/static/admin/")

    doLast {
        File("../admin/dist/").deleteRecursively()
    }
}

tasks.getByName("processResources") {
    mustRunAfter("buildAdminToBlog")
}

tasks.create(name = "buildBlog") {
    group = "build"
    dependsOn("buildAdminToBlog", "bootJar")

    doLast {
        File("./src/main/resources/static/admin/assets").deleteRecursively()
        File("./src/main/resources/static/admin/favicon.ico").delete()
        File("./src/main/resources/static/admin/index.html").delete()
    }
}
