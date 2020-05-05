window.onload = function(){
    var cron = document.getElementById("cron")

    document.querySelector(".cronexpression").addEventListener('click', async () => {
        console.log(cron.value)
        await cronexpression(cron.value); // Call Go function
    })
}