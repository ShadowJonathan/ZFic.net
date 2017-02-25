window.setTimeout(function() {
    ani.removeClass('animate');
    ani.removeClass('loading');
    ani.addClass('error');
    //noinspection JSJQueryEfficiency
    $('#ani-over').removeClass('ani-over');
    //noinspection JSJQueryEfficiency
    $('#ani-over').addClass('error');
}, 7500);

window.setTimeout(function() {
    $('#username').removeClass('minimal');
    $('#password').removeClass('minimal');
    //noinspection JSJQueryEfficiency
    $('#login-window').addClass('error');
    //noinspection JSJQueryEfficiency
    $('#ani-over').removeClass('error');
    //noinspection JSJQueryEfficiency
    $('#ani-over').addClass('errordone');
    //noinspection JSJQueryEfficiency
    $('#login-window').removeClass('submitted');
    ani.removeClass('error');
}, 10500);

window.setTimeout(function() {
    $('#ani-over').removeClass('errordone');
}, 11250);