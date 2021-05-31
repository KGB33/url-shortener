const BASE_URL = "http://127.0.0.1:8080/api/v1"

const app = new Vue({
	el: "#main",
	mounted:function() {
		this.fetchAllUrls()
	},
	data: {
		urls: [],
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
					this.urls = data;
					this.resultAvailable = true;
				})
				.catch(err => {
					console.log(err);
				});
			},
		deleteUrl(shortUrl) {
			fetch(BASE_URL + `/d/${shortUrl}`, {"method": "DELETE"})
			this.fetchAllUrls()
			},
		editUrl(shortUrl) {
			alert("Editing a URL is WIP, delete and re-create in the meantime")
			},

		}
})

const createUrlBox = new Vue({
	el: "#createBox",
	data: {
		shortUrl: "",
		destUrl: "",
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
					app.fetchAllUrls()
				}
			})
			.catch(err => {
				console.log(err);
			});

		}
	}

})
