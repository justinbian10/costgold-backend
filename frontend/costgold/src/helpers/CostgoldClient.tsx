class CostgoldClient {
	baseURL: string	

	constructor(baseURL: string) {
		this.baseURL = baseURL;
	}

	async getPrice(product: string) {
		const url = new URL(`${this.baseURL}/price/${product}`);
		const response = await fetch(url, {method: 'GET'});
		return await this._handleResponse(response);
	}

	async _handleResponse(response) {
		if (!response.ok) {
			const errorData = await response.json();
			throw new Error(errorData.message || 'Something went wrong');
		}
		return response.json();
	}
}

export default CostgoldClient;
