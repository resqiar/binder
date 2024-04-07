const editor = ace.edit("editor");
const hidden = document.getElementById("hidden-editor");
const editorLenText = document.getElementById("editor-len");
const editorMaxLen = 10000;

editor.setTheme("ace/theme/terminal");
editor.session.setMode("ace/mode/typescript");
editor.session.on("change", () => editorLenText.innerText = editor.getValue().length);
editor.session.on("change", editorChangeDebounce(function() {
    const editorValue = editor.getValue();

    // if exceeding the maximum characters -> trim the value in editor and input
    if (editorValue.length > editorMaxLen) {
        const trimmed = editorValue.substring(0, editorMaxLen);
        editor.setValue(trimmed);
        editorValue = trimmed;
    }

    hidden.value = editorValue;
}, 500))

function editorChangeDebounce(func, timeout) {
    let timer;
    return () => {
        clearTimeout(timer);
        timer = setTimeout(() => { func.apply(this) }, timeout);
    }
}
