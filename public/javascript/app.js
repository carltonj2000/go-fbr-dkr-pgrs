(() => {
  const answerWrapper = document.querySelectorAll(".answer-wrapper");
  const toggleBtns = document.querySelectorAll(".answer-toggle");

  for (const ans of answerWrapper) {
    ans.style.display = "none";
  }

  for (const btn of toggleBtns) {
    btn.addEventListener("click", (e) => {
      const answer = e.target.parentElement.nextElementSibling;
      answer.style.display = answer.style.display === "none" ? "block" : "none";
    });
  }

  const ef = document.getElementById("form-update-fact");
  const factToEdit = ef && ef.dataset.factid;

  ef &&
    ef.addEventListener("submit", (e) => {
      e.preventDefault();
      const formData = Object.fromEntries(new FormData(ef));

      return fetch(`/fact/${factToEdit}`, {
        method: "PATCH",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
      }).then(() => (document.location.href = `/fact/${factToEdit}`));
    });

  const del = document.getElementById("delete");
  const factToDelete = del && del.dataset.factid;

  del &&
    del.addEventListener("click", () => {
      const conf = confirm("Are you sure you want to delete.");
      if (!conf) return;
      return fetch(`/fact/${factToDelete}`, { method: "DELETE" }).then(
        () => (document.location.href = `/`)
      );
    });
})();
