$(function () {
    $.ajax({
            url: "http://localhost:8080/allBlock",
            data: {},
            type: "get",
            success: function f(res) {
                initBlock(res)
            },
        }
    )
})

function initBlock(res) {
    let li = document.getElementById("block_list")
    console.log(res)
    let list = JSON.parse(res)
    console.log(list)
    for(let i = 0; i < list.length; i++) {
        let id = list[i].Header.Index
        if (id == null) id = 0
        li.innerHTML +=
            "<li class='item'>" + id
            + "\t" + `<a href='singleBlock.html?hash=${list[i].Hash}'>${list[i].Hash}</a>` + "</li>";
    }
}
