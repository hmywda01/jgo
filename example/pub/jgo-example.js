JGoExampleApp = {};

$(function() {

    /**
     * @param token {string}
     */
    function InitCsrf(token) {
        /**
         * @param method {string}
         * @returns {boolean}
         */
        function csrfSafeMethod(method) {
            // HTTP methods that do not require CSRF protection.
            return (/^(GET|HEAD|OPTIONS|TRACE)$/.test(method));
        }

        $.ajaxSetup({
            crossDomain: false,
            beforeSend: function (xhr, settings) {
                if (!csrfSafeMethod(settings.type)) {
                    xhr.setRequestHeader("X-CSRF-Token", token);
                }
            }
        });
    }

    JGoExampleApp.InitCsrf = InitCsrf;

    var BaseURL = "/";

    /**
     * @param url {string}
     */
    function SetBaseUrl(url) {
        BaseURL = url;
    }

    JGoExampleApp.SetBaseUrl = SetBaseUrl;

    JGoExampleApp.Form = {
        /**
         * @param {jQuery} $ele
         */
        Signup: function($ele) {
            $ele.submit(function (e) {
                e.preventDefault();
                var username = $ele.find("[name=username]").val();
                var password = $ele.find("[name=password]").val();

                if (username.length == 0) {
                    alert("Must enter a username.");
                    return;
                }

                if (password.length == 0) {
                    alert("Must enter a password.");
                    return;
                }

                $.ajax({
                    type: "POST",
                    url: BaseURL + JGoExampleApp.URL.SignupSubmit,
                    data: {
                        username: username,
                        password: password
                    },
                    success: function () {
                        window.location = BaseURL + JGoExampleApp.URL.Lobby
                    },
                    /**
                     * @param {XMLHttpRequest} xhr
                     */
                    error: function (xhr) {
                        alert("Error creating account:\n" + xhr.responseText + "\nIf this problem persists, try refreshing the page.");
                    }
                });
            });
        },
        /**
         * @param {jQuery} $ele
         */
        Login: function($ele) {
            $ele.submit(function (e) {
                e.preventDefault();
                var username = $ele.find("[name=username]").val();
                var password = $ele.find("[name=password]").val();

                if (username.length == 0) {
                    alert("Must enter a username.");
                    return;
                }

                if (password.length == 0) {
                    alert("Must enter a password.");
                    return;
                }

                $.ajax({
                    type: "POST",
                    url: BaseURL + JGoExampleApp.URL.LoginSubmit,
                    data: {
                        username: username,
                        password: password
                    },
                    success: function () {
                        window.location = BaseURL + JGoExampleApp.URL.Lobby
                    },
                    /**
                     * @param {XMLHttpRequest} xhr
                     */
                    error: function (xhr) {
                        alert("Error logging in:\n" + xhr.responseText + "\nIf this problem persists, try refreshing the page.");
                    }
                });
            });
        }
    };

    JGoExampleApp.URL = {
        Lobby: "lobby",
        LoginSubmit: "login-submit",
        SignupSubmit: "signup-submit"
    };

});