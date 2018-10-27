const $form = document.getElementById("post");
const $name = document.getElementById("name");
const $result = document.getElementById("result");
const $metadata = document.getElementById("metadata");

$form.addEventListener("submit", (event) => {
    event.preventDefault();

    fetch("/api/v1/greet", {
        method: 'POST',
        body: JSON.stringify({
            name: $name.value
        })
    }).then(response => response.json())
      .then(json => {
          $result.innerText = json.message;
          $metadata.innerText = `The name ${json.__metadata.name} has been greeted ${json.__metadata.count} time(s)!`;
      })
      .catch(error => {
          console.error(error);
          $result.innerText = "Unexpected error!";
          $metadata.innerHTML = "";
      });
});