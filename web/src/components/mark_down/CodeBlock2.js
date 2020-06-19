import React, {PureComponent} from "react";
import PropTypes from "prop-types";
import {PrismLight as SyntaxHighlighter} from "react-syntax-highlighter";
// 设置高亮样式
import {coy} from "react-syntax-highlighter/dist/esm/styles/prism";
// 设置高亮的语言
import {
    jsx,
    javascript,
    sass,
    scss,
    c,
    cpp,
    go,
    python,
    rust,
    ruby,
    java,
    json,
    yaml,
    aspnet,
    vbnet,
    docker,
    kotlin
} from "react-syntax-highlighter/dist/esm/languages/prism";

class CodeBlock2 extends PureComponent {
    static propTypes = {
        value: PropTypes.string.isRequired,
        language: PropTypes.string
    };

    static defaultProps = {
        language: null
    };

    componentWillMount() {
        // 注册要高亮的语法，
        // 注意：如果不设置打包后供第三方使用是不起作用的
        SyntaxHighlighter.registerLanguage("jsx", jsx);
        SyntaxHighlighter.registerLanguage("javascript", javascript);
        SyntaxHighlighter.registerLanguage("sass", sass);
        SyntaxHighlighter.registerLanguage("scss", scss);
        SyntaxHighlighter.registerLanguage("c", c);
        SyntaxHighlighter.registerLanguage("cpp", cpp);
        SyntaxHighlighter.registerLanguage("go", go);
        SyntaxHighlighter.registerLanguage("python", python);
        SyntaxHighlighter.registerLanguage("rust", rust);
        SyntaxHighlighter.registerLanguage("ruby", ruby);
        SyntaxHighlighter.registerLanguage("java", java);
        SyntaxHighlighter.registerLanguage("json", json);
        SyntaxHighlighter.registerLanguage("yaml", yaml);
        SyntaxHighlighter.registerLanguage("aspnet", aspnet);
        SyntaxHighlighter.registerLanguage("vbnet", vbnet);
        SyntaxHighlighter.registerLanguage("docker", docker);
        SyntaxHighlighter.registerLanguage("kotlin", kotlin);
    }

    render() {
        const {language, value} = this.props;
        return (
            <figure className="highlight">
                <SyntaxHighlighter language={language} style={coy}>
                    {value}
                </SyntaxHighlighter>
            </figure>
        );
    }
}

export default CodeBlock2;