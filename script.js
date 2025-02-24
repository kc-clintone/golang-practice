let currentQuiz = 0;
let score = 0;
let timer;
let timeLeft;
let quizData = [];

async function fetchQuizData() {
    const response = await fetch('https://raw.githubusercontent.com/yourusername/yourrepo/main/quizData.json');
    quizData = await response.json();
    loadQuiz();
}

function loadQuiz() {
    clearInterval(timer);
    timeLeft = quizData[currentQuiz].time;
    document.getElementById('timer').innerText = `Time left: ${timeLeft}s`;
    document.getElementById('next').disabled = true;

    const quizContainer = document.getElementById('quiz');
    const currentQuizData = quizData[currentQuiz];

    quizContainer.innerHTML = `
        <h2>${currentQuizData.question}</h2>
        <label><input type="radio" name="answer" value="a"> ${currentQuizData.a}</label><br>
        <label><input type="radio" name="answer" value="b"> ${currentQuizData.b}</label><br>
        <label><input type="radio" name="answer" value="c"> ${currentQuizData.c}</label><br>
        <label><input type="radio" name="answer" value="d"> ${currentQuizData.d}</label><br>
    `;

    timer = setInterval(() => {
        timeLeft--;
        document.getElementById('timer').innerText = `Time left: ${timeLeft}s`;
        if (timeLeft <= 0) {
            clearInterval(timer);
            document.getElementById('next').disabled = false;
            alert("Time's up! Please proceed to the next question.");
        }
    }, 1000);
}

function getSelected() {
    const answers = document.querySelectorAll('input[name="answer"]');
    let answer;
    answers.forEach((ans) => {
        if (ans.checked) {
            answer = ans.value;
        }
    });
    return answer;
}

document.getElementById('next').addEventListener('click', () => {
    const answer = getSelected();
    if (answer) {
        if (answer === quizData[currentQuiz].correct) {
            score++;
        }
        currentQuiz++;
        if (currentQuiz < quizData.length) {
            loadQuiz();
        } else {
            clearInterval(timer);
            document.getElementById('quiz').innerHTML = `<h2>You scored ${score} out of ${quizData.length}</h2>`;
            document.getElementById('timer').innerText = '';
            document.getElementById('next').style.display = 'none';
        }
    }
});

fetchQuizData();
