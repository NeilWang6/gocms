"use strict";

function _defineProperty(e, t, o) {
    return t in e ? Object.defineProperty(e, t, {
        value: o,
        enumerable: !0,
        configurable: !0,
        writable: !0
    }) : e[t] = o, e
}

function headerScroll() {
    0 < $(document).scrollTop() && ($("#header").css({
        background: "rgba(0,0,0,.6)",
        height: "68px",
        "box-shadow": "0 2px 6px rgba(0, 0, 0, .1)"
    }), $("#header .header-wrapper").css({
        height: "68px",
        "border-bottom": "none"
    })), $(document).scroll(function () {
        0 < $(this).scrollTop() ? ($("#header").css({
            background: "#fff",
            height: "68px",
            "box-shadow": "0 2px 6px rgba(0, 0, 0, .1)"
        }), $("#header .header-wrapper").css({
            height: "68px",
            "border-bottom": "none"
        }), $("#header .nav").css({
            color: "#444"
        }), $("#header .header-wrapper .en").css({
            color: "#444"
        }), $("#header .logo").addClass("scroll")) : ($("#header").css({
            background: "transparent",
            height: "98px",
            "box-shadow": "none"
        }), $("#header .header-wrapper").css({
            height: "98px",
            "border-bottom": "1px solid rgba(255, 255, 255, 0.2)"
        }), $("#header .nav").css({
            color: "#fff"
        }), $("#header .header-wrapper .en").css({
            color: "#fff"
        }), $("#header .logo").removeClass("scroll"))
    }), $("#header .nav .nav-item"), $("#header .nav .move-bar").css({})
}

function navMoveBar() {
    var e = $("#header .nav .active").width(),
        t = $("#header .nav .active").position().left + parseFloat($("#header .nav .active").css("padding-left"));
    $("#header .nav .nav-item").hover(function () {
        var e, t;
        $("#header .nav .move-bar").css({
            width: (t = $(this), t.width()),
            left: (e = $(this), e.position().left + parseFloat($("#header .nav .active").css("padding-left")))
        }), $("#header .nav .move-bar").css({
            transition: ".5s"
        })
    }, function () {
        $("#header .nav .move-bar").css({
            width: e,
            left: t
        })
    }), $("#header .nav .move-bar").css({
        width: e,
        left: t
    })
}

function bannerOwl() {
    if ($("#banner .owl-carousel").length === 0) return;
    var owl = $("#banner .owl-carousel").owlCarousel({
        items: 1,
        dots: true,
        loop: true,
        nav: true,
        navText: ["", ""],
        smartSpeed: 600,
        autoplayHoverPause: true,
        autoplay: true
    });
    var autoSpeed = $("#banner .owl-carousel").data("owl.carousel").options
        .autoplayTimeout;
    var bannerNum = $("#banner .owl-carousel .owl-item:not(.cloned)").length;
    //按钮数字
    $("#banner .owl-carousel .owl-nav button").html(
        `<span>01/0${bannerNum}</span>`
    );

    //进度条
    var progressW = {
        width: '100%'
    }

    function progress(bl) {
        if (bl) {
            $('#banner .progress-bar .bar-scroll').animate(progressW, autoSpeed, function () {
                $("#banner .progress-bar .bar").css({
                    width: '0'
                })
                progress(true)
            });
        }
    }
    progress(true);

    $('#banner').hover(
        function () {
            $('#banner .progress-bar .bar-scroll').stop();
            owl.trigger('stop.owl.autoplay');
            $("#banner .progress-bar .bar").css({
                transition: 'all 0.3s ease-out 0s',
                opacity: '0'
            })
        },
        function () {
            $("#banner .progress-bar .bar").css({
                width: '0',
                transition: 'none',
                opacity: 1
            })
            progress(true);
            owl.trigger('play.owl.autoplay');
        }
    )

    $("#banner .owl-carousel").on("changed.owl.carousel", function (e) {
        $("#banner .progress-bar .bar").animate({
            width: "0"
        },
            0
        );

        $("#banner .owl-carousel .owl-nav button span").html(
            `0${e.page.index + 1}/0${bannerNum}`
        );
    });
}

function groupParallax() {
    $(document).scroll(function () {
        var e = .2 * $(document).scrollTop();
        252 < $(document).scrollTop() && $("#group-intr").css({
            backgroundPositionY: -e
        })
    })
}

function newOwl() {
    0 === $("#news .owl-carousel").length && $(document).scrollTop() < 2030 || $("#news .owl-carousel").owlCarousel({
        items: 3,
        dots: !1,
        loop: !0,
        nav: !0,
        navText: ["", ""],
        smartSpeed: 600,
        autoplay: !0
    })
}

function logoGroup() {
    var e = function (o) {
        $("#invest .invest-list").find("[data-id]").each(function (e, t) {
            $(t).data("id") === o && $(t).addClass("show")
        })
    },
        t = function (e) {
            $("#invest .invest-list").find("[data-id]").each(function (e, t) {
                $(t).removeClass("show")
            })
        };
    $("#invest").find("[data-id]").mousemove(function () {
        t(), e($(this).data("id"))
    }), $("#invest .invest-list").mouseleave(function () {
        t()
    })
}

function count() {
    if (0 !== $("#company-intr .data-item").length) {
        var e = {
            useEasing: !0,
            useGrouping: !0,
            separator: ""
        },
            t = new CountUp("count1", 0, $("#count1").text(), 0, 2.5, e),
            o = new CountUp("count2", 0, $("#count2").text(), 0, 2.5, e),
            n = new CountUp("count3", 0, $("#count3").text(), 0, 2.5, e),
            a = new CountUp("count4", 0, $("#count4").text(), 0, 2.5, e),
            s = new CountUp("count5", 0, $("#count5").text(), 0, 2.5, e),
            r = new CountUp("count6", 0, $("#count6").text(), 0, 2.5, e);
        $(document).scroll(function () {
            2500 < $(this).scrollTop() && (t.start(), o.start(), n.start(), a.start(), s.start(), r.start())
        })
    }
}

function countT() {
    if (0 !== $("#teach-data .data-list li").length) {
        var e = {
            useEasing: !0,
            useGrouping: !0,
            separator: ""
        },
            t = new CountUp("data1", 0, $("#data1").text(), 0, 2.5, e),
            o = new CountUp("data2", 0, $("#data2").text(), 0, 2.5, e),
            n = new CountUp("data3", 0, $("#data3").text(), 0, 2.5, e),
            a = new CountUp("data4", 0, $("#data4").text(), 0, 2.5, e),
            s = new CountUp("data5", 0, $("#data5").text(), 0, 2.5, e),
            r = new CountUp("data6", 0, $("#data6").text(), 0, 2.5, e);
        t.start(), o.start(), n.start(), a.start(), s.start(), r.start()
    }
}

function honorSly() {
    var itemLeft = 0;
    $("#honor .honor-list .honor-con").each(function (i, n) {
        if (i - 1 >= 0) {
            itemLeft += parseFloat($("#honor .honor-list .honor-con").eq(i).css('width'));
        }
        $(n).css({
            left: itemLeft
        })
    })

    var move = false;
    var moveIn = false;
    var goPos = 0;
    var posX = 0;
    var max = 0;
    var min = -parseInt($('#honor .sly-wrapper .honor-con:last').css('left'));

    $('#honor .sly-wrapper').mousemove(
        function (ev) {
            var _width = parseInt($(this).css('width')) / 2;
            posX = ev.clientX;
            moveIn = true;
            if (!move) {
                move = true;

                var list = $('.honor-list', this);

                function go() {
                    requestAnimationFrame(function () {
                        if (!moveIn) return false;
                        if (goPos < min) {
                            goPos = min
                            return false;
                        } else if (goPos > max) {
                            goPos = max
                            return false;
                        };
                        go();

                        if (posX > _width) {

                            goPos -= 1;
                        } else {

                            goPos += 1;

                        }
                        list.css({
                            left: "" + (goPos) + "px"
                        })
                        $("#honor .mark-wrapper").css({
                            left: goPos * 1.3
                        });
                    })
                }

                go();
            }
        }
    )
    $('#honor .sly-wrapper').mouseleave(function () {
        move = false;
        moveIn = false;
    })

}

function aside() {
    function num(n) {
        return $(`#scope .scope-aside .aside-list ul:eq(${n}) li`).length
    }

    function show(id) {
        $('#scope .scope-content .content-wrapper').each(function (i, n) {
            $(this).removeClass('show');
            if ($(this).data('id') === id) {
                $(this).addClass('show');
            }
        })

    }

    $("#scope .scope-aside .aside-list:eq(0) ul").css({
        height: num($(this).find('li').length) * 50 + 'px'
    })

    //进入页面后显示
    $("#scope .scope-aside .aside-list:eq(0) li:eq(0)").addClass("select");
    show(0)

    $("#scope .scope-aside .aside-list").click(function () {
        $("#scope .scope-aside .aside-list").each(function () {
            $(this).find('ul').css({
                'height': 0
            });
        });
        $(this).find('ul').css({
            'height': num($(this).index()) * 50 + 'px'
        });
    });
    $("#scope .scope-aside .aside-list li").click(function () {
        $("#scope .scope-aside .aside-list li").each(function () {
            $(this).removeClass("select");
        });
        $(this).addClass("select");
        show($(this).data('id'))
    });
}

function wow() {
    $('#content-list li').addClass('wow fadeInUp');
    $('#content-list li').each(function (i, n) {
        $(n).attr('data-wow-delay', `${((i + 1) * 0.2).toFixed(1)}s`);
    })

    $('#invest .invest-list .list-logo li').addClass('wow zoomIn');
    $('#invest .invest-list .list-logo li').each(function (i, n) {
        $(n).attr('data-wow-delay', `${((i % 5 + 1) * 0.2).toFixed(1)}s`);
    })

    $('#news-list .list-wrapper .news-con-item').addClass('wow fadeInUp');
    $('#news-list .list-wrapper .news-con-item').each(function (i, n) {
        $(n).attr('data-wow-delay', `${((i % 3 + 1) * 0.2).toFixed(1)}s`);
    })

    $('#news .active .news-con-item').addClass('wow fadeInUp');
    $('#news .active .news-con-item').each(function (i, n) {
        $(n).attr('data-wow-delay', `${((i + 1) * 0.2).toFixed(1)}s`);
    })

    new WOW().init();
}

// 数组去重
function unique(arr) {
    if (!Array.isArray(arr)) {
        console.log('type error!')
        return
    }
    let res = []
    for (let i = 0; i < arr.length; i++) {
        if (res.indexOf(arr[i]) === -1) {
            res.push(arr[i])
        }
    }
    return res
}

function moveUlBar(key) {
    // var ul = $(the).children("ul");
    // ul.height(120);
    $(".aside-list").each(function (i,o) {
        if ($(o).children("ul").hasClass("ul-"+key)) {
            $(o).children("ul").height(200);
        }else {
            $(o).children("ul").height(0);
        }
    })
}

function moveLiSelect(obj,key) {
    //console.log($(obj).parent("ul"));
    // 菜单选中
    var par = $(obj).parent("ul");
    $(par).children("li").each(function (i,o) {
        if ($(o).attr("data-id") == key) {
            $(o).addClass("select");
        }else {

            $(o).removeClass("select");
        }
    });
    // 内容选中
    $(".content-wrapper").each(function (i,o) {
        if ($(o).attr("data-id") == key) {
            $(o).addClass("show");
        }else {
            $(o).removeClass("show");
        }
    })
}

function insertUnit() {
    $(".aside-list").onclick = function () {
        alert(1);
    }
    // if ($('#scope').length === 0) return;
    // var data = JSON.parse($(".getData_unit").text());
    // var newArr = [];
    // var newData = [];
    //
    // $.each(data, function (i, n) {
    //     n.id = i;
    // })
    //
    // $.each(data, function (i, n) {
    //     newArr.push((n.classify).trim());
    // })
    //
    // $.each(unique(newArr), function (i, n) {
    //     var obj = {};
    //     obj.classify = n;
    //     obj.list = [];
    //     newData.push(obj)
    // })
    //
    // $.each(data, function (i, n) {
    //     $.each(newData, function (j, k) {
    //         if ((n.classify.trim()) === k.classify) {
    //             // delete n.classify
    //             k.list.push(n);
    //         }
    //     })
    // })
    //
    // var $parentAside = $('#scope .scope-aside');
    // var strAside = '';
    // var strLi = '';
    //
    //
    // $.each(newData, function (i, n) {
    //     $.each(n.list, function (j, k) {
    //         strLi += `<li data-id="${k.id}">${k.title}</li>`;
    //     })
    //     strAside += `<div class="aside-list"><h3>${n.classify}</h3> <ul>${strLi}</ul></div>`;
    //     strLi = '';
    // })
    // $parentAside.html(strAside);
    //
    // aside();

}

//新闻详情
function newsDetail() {
    $('#news-wrapper .news-btn a').each(function () {
        if (!$(this).text()) {
            $(this).css({
                background: '#f9f9f9',
                cursor: 'not-allowed'
            });
            $(this).attr('href', 'javascript:;');
        }
    })
}

function isPC() {
    var userAgentInfo = navigator.userAgent;
    var Agents = ["Android", "iPhone",
        "SymbianOS", "Windows Phone",
        "iPad", "iPod"];
    var flag = true;
    for (var v = 0; v < Agents.length; v++) {
        if (userAgentInfo.indexOf(Agents[v]) > 0) {
            flag = false;
            break;
        }
    }
    if (flag === false) {
        var url = document.location.host;
        var href = document.location.href;
        if (href.match("/m/")) {
            return
        }
        var newUrl = href.replace(url, url + "/m");
        console.log(newUrl);
        window.location.href = newUrl;
        return
    }
    return flag;
}

$(function () {
    headerScroll(), navMoveBar(), bannerOwl(), groupParallax(), newOwl(), logoGroup(), count(), honorSly(), countT(), wow(), insertUnit(), newsDetail()
});