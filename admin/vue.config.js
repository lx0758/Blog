module.exports = {
    publicPath: '/admin/',
    outputDir: '../blog/src/main/resources/static/admin/',
    devServer: {
        port: 8081,
        proxy: {
            '^/api/.*': {
                target: 'http://localhost:8080',
                ws: true,
                secure: false,
                changeOrigin: true,
            },
            '^/vcode': {
                target: 'http://localhost:8080',
                ws: true,
                secure: false,
                changeOrigin: true,
            },
        },
    },
}