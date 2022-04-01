import org.jetbrains.kotlin.gradle.tasks.KotlinCompile

plugins {
    id("org.springframework.boot") version "2.6.6"
    id("io.spring.dependency-management") version "1.0.11.RELEASE"
    kotlin("jvm") version "1.6.10"
    kotlin("plugin.spring") version "1.6.10"
    id("org.jetbrains.kotlin.plugin.noarg") version "1.6.10"
}

group = "com.liux.blog"
version = "0.0.1-SNAPSHOT"
java.sourceCompatibility = JavaVersion.VERSION_1_8

dependencies {
    implementation("org.jetbrains.kotlin:kotlin-reflect")
    implementation("org.jetbrains.kotlin:kotlin-stdlib-jdk8")

    implementation("org.springframework.boot:spring-boot-starter")
    implementation("org.springframework.boot:spring-boot-starter-web")
    implementation("org.springframework.boot:spring-boot-starter-websocket")
    implementation("org.springframework.boot:spring-boot-starter-thymeleaf")
    implementation("org.springframework.boot:spring-boot-starter-security")
    implementation("org.springframework.boot:spring-boot-starter-data-redis")
    implementation("org.springframework.boot:spring-boot-starter-mail")

    implementation("org.springframework.session:spring-session-data-redis")

    implementation("com.github.pagehelper:pagehelper-spring-boot-starter:1.4.1")
    implementation("org.mybatis.spring.boot:mybatis-spring-boot-starter:2.2.2")
    implementation("com.fasterxml.jackson.module:jackson-module-kotlin:2.13.2")
    implementation("com.github.xiaoymin:knife4j-spring-ui:3.0.3")
    implementation("io.springfox:springfox-swagger2:3.0.0")
    implementation("com.github.ua-parser:uap-java:1.5.2")
    implementation("com.github.penggle:kaptcha:2.3.2")

    implementation("org.commonmark:commonmark:0.18.2")
    implementation("org.commonmark:commonmark-ext-autolink:0.18.2")
    implementation("org.commonmark:commonmark-ext-gfm-tables:0.18.2")
    implementation("org.commonmark:commonmark-ext-gfm-strikethrough:0.18.2")

    developmentOnly("org.springframework.boot:spring-boot-devtools")
    runtimeOnly("mysql:mysql-connector-java")

    testImplementation("org.springframework.boot:spring-boot-starter-test")
    testImplementation("org.springframework.security:spring-security-test")
}

tasks.withType<KotlinCompile> {
    kotlinOptions {
        freeCompilerArgs = listOf("-Xjsr305=strict")
        jvmTarget = "1.8"
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
        commandLine = listOf("node", "./node_modules/@vue/cli-service/bin/vue-cli-service.js", "build")
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

tasks.getByName("bootJar") {
    mustRunAfter("buildAdminToBlog")
}

tasks.create(name = "buildBlog") {
    group = "build"
    dependsOn("buildAdminToBlog", "bootJar")

    doLast {
        File("./src/main/resources/static/admin/css").deleteRecursively()
        File("./src/main/resources/static/admin/js").deleteRecursively()
        File("./src/main/resources/static/admin/fonts").deleteRecursively()
        File("./src/main/resources/static/admin/favicon.ico").delete()
        File("./src/main/resources/static/admin/index.html").delete()
    }
}
