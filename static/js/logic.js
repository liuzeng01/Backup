window.onload = function () {

    var localconfigitems = document.getElementsByName("localconfigitems");
    var sftpconfigitems = document.getElementsByName("sftpconfigitems");
    var Encryptconfigitems = document.getElementsByName("Encryptconfigitems");

    var save1 = document.getElementById("saveconfigitems");
    var localitems = new Object();
    var sftpitems = new Object();
    var Encryptitems = new Object();




    save1.addEventListener("click", async () => {

        for (var i = 0; i < localconfigitems.length; i++) {
            localitems[localconfigitems[i].title] = localconfigitems[i].value;
        }
        for (var i = 0; i < sftpconfigitems.length; i++) {
            sftpitems[sftpconfigitems[i].title] = sftpconfigitems[i].value;

        }
        for (var i = 0; i < Encryptconfigitems.length; i++) {
            Encryptitems[Encryptconfigitems[i].title] = Encryptconfigitems[i].value;
        }

        await configitemsSave(localitems, sftpitems, Encryptitems);
        location.reload()

    });

    var build = document.getElementById("buildpackage")
    build.addEventListener("click", async () => {
        console.log("Build Package !!!")
        await buildpackage()
        location.reload()
    })


    var $editor = $("#editor");
    $editor.markdown({
        height: 300,
        onShow: function () {}
    });


    





}