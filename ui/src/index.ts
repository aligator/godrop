import App from './App'

function run() {
    document.getElementById("root")?.append(...App().children)
}

run()