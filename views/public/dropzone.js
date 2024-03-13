// @ts-check

/**
* @type {HTMLInputElement | null}
*/
const dzInput = document.getElementById("dropzone");
const dzPreview = document.getElementById("dropzone-preview");
const dzRestoreElem = document.getElementById("dropzone-reset");

/**
* @type {FileList | null}
*/
let dzFiles = null;

dzRestoreElem?.addEventListener("click", dzRestore)
dzInput?.addEventListener("change", function() {
    if (!this.files) return;
    if (!dzValidate(this.files)) return dzInput.value = "";

    dzFiles = this.files;
    dzPreviews();
})

function dzPreviews() {
    if (!dzFiles) return;

    /**
    * @type {Node[]}
    */
    const children = [];

    for (let i = 0; i < dzFiles.length; i++) {
        const file = dzFiles[i];
        if (!file) continue;

        // remove all children before appending new child
        dzPreview?.replaceChildren();

        const reader = new FileReader();
        reader.onload = function(e) {
            const child = document.createElement("img");
            child.src = e.target?.result?.toString() ?? "";
            children.push(child);

            // only invoke when reading reach the last file
            if (children.length === dzFiles?.length) {
                dzPreview?.replaceChildren(...children);
                dzShowReset();
            }
        }
        reader.readAsDataURL(file);
    }
}

/**
* @param files {FileList}
* @returns boolean
*/
function dzValidate(files) {
    if (files.length > 5) {
        alert("Max number of images are only 5");
        return false;
    }

    for (let i = 0; i < files.length; i++) {
        const file = files[i];

        if (!file.type.startsWith("image/")) {
            alert(`"${file.name}" is not an image`);
            return false;
        };

        const maxSize = 1 * 1024 * 1024; // 1 MB
        if (file.size > maxSize) {
            alert(`"${file.name}" is exceeding 1MB size`);
            return false;
        }
    }

    return true;
}

function dzRestore() {
    if (!dzPreview || !dzInput) return;

    dzFiles = null;
    dzPreview.replaceChildren();
    dzInput.value = "";

    dzShowReset(false);
}

/**
* @param show {Boolean}
* @returns void
*/
function dzShowReset(show = true) {
    if (!dzRestoreElem) return;
    if (show) {
        dzRestoreElem.style.display = "flex";
    } else {
        dzRestoreElem.style.display = "none";
    }
}

const dzArea = document.getElementById("droppable-area");

;['dragenter', 'dragover', 'dragleave', 'drop'].forEach(eventName => {
    dzArea?.addEventListener(eventName, function(e) {
        e.preventDefault();
        e.stopPropagation();
    })
});

;['dragenter', 'dragover'].forEach(eventName => {
    dzArea?.addEventListener(eventName, () => dzHighlight(true))
});

;['dragleave', 'drop'].forEach(eventName => {
    dzArea?.addEventListener(eventName, () => dzHighlight(false))
});

/**
* @param h {Boolean}
* @returns void
*/
function dzHighlight(h = true) {
    if (!dzArea) return;
    h ? dzArea.style.border = "2px dashed yellow" :
        dzArea.style.border = "none";
}

dzArea?.addEventListener("drop", dzDrop);

/**
* @param e {DragEvent}
* @returns void
*/
function dzDrop(e) {
    let dataTransfer = e.dataTransfer;
    let files = dataTransfer?.files;

    if (!dzInput || !files) return;
    if (!dzValidate(files)) return dzInput.value = "";

    dzFiles = files;
    dzPreviews();

    // inject dragged files into HTML input
    dzInput.files = files;
}
