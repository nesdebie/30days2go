async function play(choice) {
  const response = await fetch("/play", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ player: choice })
  });

  const data = await response.json();
  document.getElementById("result").innerHTML =
    `You chose <b>${data.player}</b><br>` +
    `The computer chose <b>${data.computer}</b><br>` +
    `<strong>Result: ${data.result}</strong>`;
}
