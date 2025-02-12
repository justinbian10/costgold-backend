import { useState } from 'react'
import ProductForm from "./components/ProductForm.tsx";
import PurchaseValueDisplay from "./components/PurchaseValueDisplay.tsx";

import CostgoldClient from "./helpers/CostgoldClient.tsx";

const BASE_URL = 'https://costgold-backend-b740d2e4daaf.herokuapp.com';

function App() {
  const [price, setPrice] = useState(0);
  const [isPriceLoading, setIsPriceLoading] = useState(false);

  const handlePriceQuery = async (event) => {
  	const client = new CostgoldClient(BASE_URL);
  	const response = client.getPrice(event.target.product.value);
	setIsPriceLoading(true);
	const newPrice = (await response).Price;
	setIsPriceLoading(false);
	console.log(newPrice)
	setPrice(newPrice);
  }
  
  const displayPrice = () => {
	if (isPriceLoading) {
		return "Loading";
	}
	if (price === 0) {
		return ""
	}
	return `$${price.toFixed(2)}`;
  }

  return (
    <>
    	<div className="flex items-center flex-col mx-auto" >
		<ProductForm handlePriceQuery={handlePriceQuery}/>
		<p>{displayPrice()}</p>
		<PurchaseValueDisplay price={price} buyPrice={price} />
	</div>
    </>
  )
}

export default App
