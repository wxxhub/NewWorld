function pwdCheck() {

	if (pwd.value == ensure_pwd.value) {
		help_inline.style.display = "none";
		return true;
	} else {
		help_inline.style.display = "";
		return false;
	}
};

function logonCheck() {
	if (email.value == "") {
		alert("邮箱不能为空!");
		return false;
	}

	if (user_id.value == "") {
		alert("账户不能为空!");
		return false;
	}

	if (user_name.value == "") {
		alert("昵称不能为空!");
		return false;
	}

	if (pwd.value == "") {
		alert("密码不能为空!");
		return false;
	}

	if (pwd.value != ensure_pwd.value) {
		alert("确认密码和密码不同!");
		return false;
	}

	return true;
};

function logon() {
	if (logonCheck() == true) {
		var postData = {
			"user_id":user_id.value,
			"name":user_name.value,
			"pwd":hex_sha1(pwd.value),
		};
		jQuery.ajax({
			type: "POST",
			async: false,
			dataType: "json",
			url: "/logon",
			// data: $('#user_form').serialize(),
			data: postData,
			statusCode: {
				200:function(){
					alert("注册成功!");
				},
				204:function(){
					alert("用户ID已存在!");
				},
			},
			error: function(result) {
				alert("注册异常, 请稍候再试!");
			}
		});
	}
}