let previousPressed;
const myHeaders = new Headers();
myHeaders.append('pragma', 'no-cache');
myHeaders.append('cache-control', 'no-cache');

const requestInit = {
    method: 'GET',
    headers: myHeaders,
};

function toggleMenu(id) {
    const button = document.getElementById(id);
    const buttons = document.getElementById('docs-buttons');
    const article = document.getElementById('docs-article');
    if (button.innerText === "Open") {
        button.innerText = 'Close';
        buttons.style.width = '50%';
        article.style.width = '50%';
    } else {
        button.innerText = 'Open';
        buttons.style.width = '20%';
        article.style.width = '80%';
    }
}

function pressButton(button) {
    button.classList.remove("unpressed-button");
    button.classList.add("pressed-button");
}

function unpressedButton(button) {
    button.classList.remove("pressed-button");
    button.classList.add("unpressed-button");
}

function openMdArticle(id) {
    const button = document.getElementById(id);
    if (previousPressed !== undefined) {
        unpressedButton(previousPressed);
    }
    previousPressed = button;
    pressButton(button);
    fetch(`/plasma/static/articles/${id}.md`, requestInit)
        .then(
            response => response.text()
        )
        .then(
            body => {
                const article = document.getElementById('article-body');
                const conv = new showdown.Converter();
                article.innerHTML = conv.makeHtml(body);
                Prism.highlightAll();
            }
        )
    return false;
}