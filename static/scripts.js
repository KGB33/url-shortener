const BASE_URL = "http://127.0.0.1:8080/api/v1"

const app = new Vue({
	el: "#main",
	data: {
		result: "",
		resultAvailable: false,
	},
	methods: {
		fetchAllUrls() {
			this.resultAvailable = false;

			fetch(BASE_URL + "/r", {
				"method": "GET",
				"headers": {
					//"api-key": "key-go-here",
				}

				}).then(response => {
					if(response.ok) {return response.json()}
					else {
						alert("Server returned: " + response.status + " : " + response.statusText);
					}
				})
				.then(data => {
					console.log(data);
					this.result = data;
					this.resultAvailable = true;
				})
				.catch(err => {
					console.log(err);
				});
			}
		}
})

const createUrlBox = new Vue({
	el: "#createBox",
	data: {
		shortUrl: "foo",
		destUrl: "bar",
		created: false
	},
	methods: {
		createUrl() {
			fetch(BASE_URL + "/c", {
				"method": "POST",
				"body": `{"ShortUrl": "${this.shortUrl}", "DestUrl": "${this.destUrl}"}`
			}).then(response => {
				return response.json()
			}).then(data => {
				if(data.Error) {
					alert("Server returned: " + data.Error);
				} else {
					this.ShortUrl = data.ShortUrl;
					this.DestUrl = data.DestUrl;
					this.created = true
				}
			})
			.catch(err => {
				console.log(err);
			});

		}
	}

})
