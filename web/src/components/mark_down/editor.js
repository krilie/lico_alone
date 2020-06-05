import CodeMirrorEditor from "./CodeMirrorEditor"
const React = require('react')
const PropTypes = require('prop-types')

export default function Editor(props) {
    return (
        <form className="editor pure-form">
            <CodeMirrorEditor mode="markdown" theme="monokai" value={props.value} onChange={props.onChange} />
        </form>
    )
}

Editor.propTypes = {
    onChange: PropTypes.func.isRequired,
    value: PropTypes.string
}

Editor.defaultProps = {
    value: ''
}
