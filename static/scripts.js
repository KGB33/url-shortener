const BASE_URL = "http://127.0.0.1:8080"

const app = new Vue({
	el: "#main",
	data: {
		result: "",
		resultAvailable: false,
	},
	methods: {
		fetchAllUrls() {
			this.resultAvailable = false;
			console.log("Hello fetchAllUrls");

			fetch(BASE_URL, {
				"method": "GET",
				"headers": {
					//"api-key": "key-go-here",
					"Access-Control-Allow-Origin": "127.0.0.1",
				}

				}).then(response => {
					if(response.ok) {return response.json}
					else {
						alert("Server returned: " + response.status + " : " + response.statusText);
					}
				})
				.then(response => {
					this.result = response.body;
					this.resultAvailable = true;
				})
				.catch(err => {
					console.log(err);
				});
			}
		}
})
