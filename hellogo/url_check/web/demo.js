(function($) {

	'use strict';

	$(function() {



		var tgListItem = $('.tg-list');			 // ❶ 항목이 위치하는 곳
		var tgListInput = $('.tg-list-input');		// ❷ 입력 박스 : 타겟그룹 명

		$('.first-card > .btn-check').on('click', function(){
			tgListItem.empty();
			tgListInput.focus();
		})

		// 검색버튼클릭 또는 엔터키
		tgListInput.on("keydown", function(event) {
			if (event.keyCode == 13) {
				getTargetGroupList()
			}
		});
		$('.tg-list-add-btn').on("click", function(event) {
			// event.preventDefault(); // ❸ add 버튼 클릭시. child에게 click 이벤트 전파 차단
			getTargetGroupList()
		});

		$('.tg-list-empty-btn').on("click", function(event) {
			tgListItem.empty();
			tgListInput.focus();
		});

		// 타겟그룹리스트 조회
		function getTargetGroupList() {
			var profile = $(".first-card > .btn-check:checked").val();
			var tgname = tgListInput.val();

			$.post("/targetgroup",
				{
					tgname:tgname
				},
				function(items){
					addItem(items)
				}
			);
		}


		var addItem = function(item) {		// ➎ 항목 추가 함수
			tgListItem.empty();
			console.log(item.Urls);
			for (const [key, value] of Object.entries(item.Urls)) {
				tgListItem.append("<li class='m-0 p-0'><div class='m-0 p-0 form-check'><label class='m-0 p-0 form-check-label'><input class='checkbox' type='checkbox' /><span class='list-group-item-light list-group-item-action tgitem small' id='"+ key + "' style='color:" + value + ";'>" + key + "</span><div></div><i class='input-helper'></i></label></div></li>");
			}
			//item.Urls.forEach((value,key)=>{
				//tgListItem.append("<li class='m-0 p-0' style='color:" + value +";'><div class='m-0 p-0 form-check'><label class='m-0 p-0 form-check-label'><input class='checkbox' type='checkbox' /><span class='list-group-item-light list-group-item-action tgitem small' id='"+ key + "'>" + key + "</span><div></div><i class='input-helper'></i></label></div></li>");
			//});
		};
	
	// 리스트 클릭이벤트 정의
	tgListItem.on('click', '.tgitem', function() {
		var profile = $(".first-card > .btn-check:checked").val();
		var targetGroupArn = $(this).attr('id')
		// console.log($(this))
		$.ajax({
			url: "/targethealth",
			type: "POST",
			data: {
				url: targetGroupArn
			},
			//data: JSON.stringify({profile:profile, TargetGroupArn : targetGroupArn}),
			context: $(this).next(),
			success: function(items) {
				var ul = document.createElement('ul');

				//var span = document.createElement('span');
				//var loadBalancerArn = $(this).get(0).id;
				//var parts = loadBalancerArn.split('/');
				//var loadBalancerName = parts[parts.length -2];
				//if (loadBalancerName == undefined) {
					//loadBalancerName = '';
				//}
//
				//span.innerHTML = "&nbsp;&nbsp;<small>∟&nbsp;Load Balancer : " + loadBalancerName + "</small>"
				//ul.appendChild(span)

				//var span = document.createElement('span');
				//span.innerHTML = "<br>&nbsp;&nbsp;<small>∟&nbsp;TG ARN: " + targetGroupArn  +"</small>"
				//ul.appendChild(span)
//
				var span = document.createElement('span');
				span.innerHTML = "&nbsp;&nbsp;&nbsp;<small>➟ Health : " + items.Status + "</small>";
				ul.appendChild(span);

				//if (!!items && !!items.TargetHealthDescriptions) {
//
					//items.TargetHealthDescriptions.forEach(e=>{
						//var li = document.createElement('li');
						//li.className="p-0 m-0";
						//var state = "";
						//if (e.TargetHealth.State == "healthy") {
							//state = "✔️ " + e.TargetHealth.State;
						//} else if (e.TargetHealth.State == "draining") {
							//state = "💧 " + e.TargetHealth.State;
						//} else {
							//state = "❌ " + e.TargetHealth.State;
						//}
						//li.innerHTML = "&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<small>Id: "+ e.Target.Id
							//+ ",	Port: " + e.Target.Port
							//+ ",	State: " + state
							//+ "&nbsp;&nbsp;</small>";
						//ul.appendChild(li);
					//});
				//}
				$(this).empty()
				$(this).append(ul)
			},
		});
	});



	
	});
})(jQuery);
