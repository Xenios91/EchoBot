function validateURL(url) {
    if (!url.startsWith("http://") && !url.startsWith("https://")) {
        url = "http://" + url;
    }

    const pattern = new RegExp("((http|https)://)(www.)?[a-zA-Z0-9@:%._\\+~#?&//=]{2,256}\\.[a-z]{2,6}\\b([-a-zA-Z0-9@:%._\\+~#?&//=]*)");

    if (!pattern.test(url)) {
        document.getElementById("url-error").style.display = "block";
        return false;
    } else {
        document.getElementById("url-error").style.display = "none";
        return true;
    }
}

function resetForm() {
    document.getElementById("url").classList.remove("is-invalid");
}

function setInvalidInput() {
    document.getElementById("url").classList.add("is-invalid");
}

function methodSelectionChange() {
    const requestBody = document.getElementById("request-body");
    if (document.getElementById("http-method").value === "GET") {
        requestBody.style.display = "none";
    } else {
        requestBody.style.display = "block";
    }
}

$("#apiForm").submit(function (e) {

    e.preventDefault();
    resetForm();

    var form = $(this);
    var url = form.attr("action");

    if (!validateURL(document.getElementById("url").value)) {
        return;
    }

    $.ajax({
        type: "POST",
        url,
        data: form.serialize(),
        success: (data) => {
            const object = JSON.parse(data);
            document.getElementById("status-code").value = object.StatusCode;
            document.getElementById("resp-content-type").value = object.ContentType;
            document.getElementById("response").value = object.Body;
        },
        error: () => {
            alert("An unknown error has occurred");
        }
    });
});
