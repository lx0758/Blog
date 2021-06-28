module.exports = {
    publicPath: '/admin/',
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