import HtmlWebpackPlugin from 'html-webpack-plugin';

const config = {
    mode: 'development',
    entry: './src/index.tsx',
    devServer: {
      contentBase: './dist',
    },
    output: {
        filename: 'bundle.js',
        path: new URL('./dist', import.meta.url).pathname,
        clean: true,
    },
    plugins: [
        new HtmlWebpackPlugin({
            title: 'Commit Logs',
            template: 'src/index.html',
            favicon: "src/favicon.ico"
        }),
    ],
    module: {
        rules: [
            {
                test: /\.tsx?$/,
                use: 'ts-loader',
                exclude: /node_modules/,
            },
            {
                test: /\.scss$/i,
                use: ['style-loader', 'css-loader', 'sass-loader'],
            },
            {
                test: /\.(png|svg|jpg|jpeg|gif)$/i,
                type: 'asset/resource',
            },
        ],
    },
    resolve: {
        extensions: ['.tsx', '.ts', '.js'],
    },
}

export default config;