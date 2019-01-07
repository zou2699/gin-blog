var str_svg1 = "                <svg xmlns='http://www.w3.org/2000/svg' width='24' height='24' viewBox='0 0 24 24' fill='none'\n" +
    "                     stroke='currentColor' stroke-width='2' stroke-linecap='round' stroke-linejoin='round'\n" +
    "                     class='feather feather-calendar'>\n" +
    "                    <rect x='3' y='4' width='18' height='18' rx='2' ry='2'></rect>\n" +
    "                    <line x1='16' y1='2' x2='16' y2='6'></line>\n" +
    "                    <line x1='8' y1='2' x2='8' y2='6'></line>\n" +
    "                    <line x1='3' y1='10' x2='21' y2='10'></line>\n" +
    "                </svg>"

var str_svg2 = "              <svg xmlns='http://www.w3.org/2000/svg' width='24' height='24' viewBox='0 0 24 24' fill='none'\n" +
    "                     stroke='currentColor' stroke-width='2' stroke-linecap='round' stroke-linejoin='round'\n" +
    "                     class='feather feather-tag'>\n" +
    "                    <path d='M20.59 13.41l-7.17 7.17a2 2 0 0 1-2.83 0L2 12V2h10l8.59 8.59a2 2 0 0 1 0 2.82z'></path>\n" +
    "                    <line x1='7' y1='7' x2='7' y2='7'></line>\n" +
    "                </svg>"

$(document).ready(function () {
    $.getJSON("/api/v1/articles", function (data) {
        $.each(data["data"]["list"], function () {
            console.log(this["id"])
            // $("#qaz").attr("href","/api/v1/articles/"+this["id"])
            // $("#qaz").text(this["title"])
            //
            // $("#qwe").attr("href","/api/v1/tags/"+this["tag"]["id"])
            // $("#qwe").text(this["tag"]["name"])
            var d = formatDate(this['UpdatedAt'])
            var str_article = "<a class='font-125' href='/api/v1/articles/" + this['id'] + "'>" + this['title'] + "</a>";
            var str_date = "<time datetime='date'>" + d + "</time>";
            var str_tag = "<a class='btn btn-sm btn-outline-dark tag-btn' href='/api/v1/tags/" + this['tag']['id'] + "'>" + this['tag']['name'] + "</a>";


            var a = "<p>" + str_article + "<br>" + str_svg1 + str_date + "<br>" +str_svg2+ str_tag + "</p>";
            //a.appendTo($("#main")
            $("#main").append(a)
            console.log(a)
        });
    });
})


function formatDate(date) {
    var d = new Date(date),
        month = '' + (d.getMonth() + 1),
        day = '' + d.getDate(),
        year = d.getFullYear();

    if (month.length < 2) month = '0' + month;
    if (day.length < 2) day = '0' + day;

    return [year, month, day].join('-');
}