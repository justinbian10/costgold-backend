import { useState, useEffect } from 'react'

function PurchaseValueDisplay({price, buyPrice}) {
	const [pureStatusPercent, setPureStatusPercent] = useState(.75);

	const calculateBreakeven = (feePercent) => {
		const pureFee = 0.01 * parseFloat(pureStatusPercent);
		const priceAfterPure = price - (price * pureFee)
		console.log(priceAfterPure)
		const maxBuyPrice = priceAfterPure / (1 + (0.01 * feePercent))
		return `$${(maxBuyPrice).toFixed(2)}`;
	}

	const calculatePercentage = () => {

	}

	const handleStatusChange = (event) => {
		setPureStatusPercent(event.target.value);
	}

	const displayPaymentSummary = (paymentName, feePercent) => {
		return (
			<>
				<h3 className="col-span-2">{paymentName}</h3>
				<p className="border text-center w-full">{calculateBreakeven(feePercent)}</p>
				<p className="text-center">{calculatePercentage()}</p>
			</>
		)
	}

	return (
		<div className="grid grid-cols-4 gap-y-2 p-10 w-full">
			<PaymentSummary price={price} paymentName="Kasheesh Full Fee*" feePercent="2" pureStatusPercent={pureStatusPercent} />
			<PaymentSummary price={price} paymentName="Citi AA Biz + PBP*" feePercent="-1.8" pureStatusPercent={pureStatusPercent} />
			<PaymentSummary price={price} paymentName="Venmo + PBP" feePercent="-5" pureStatusPercent={pureStatusPercent} />
			<p className="col-span-2">Pure Status</p>
			<select name="product" className="border text-md appearance-none rounded-sm p-1 col-span-2" onChange={handleStatusChange}>
				<option value=".75">Copper</option>
				<option value=".7">Silver</option>
				<option value=".625">Gold</option>
				<option value=".5">Plum</option>
			</select>

		</div>
	)

			//{displayPaymentSummary("Kasheesh Full Fee*", 2)}
			//{displayPaymentSummary("Citi AA Biz + PBP*", -1.8)}
			//{displayPaymentSummary("Venmo + PBP", -5)}
}


function PaymentSummary({price, paymentName, feePercent, pureStatusPercent}) {
	const [breakeven, setBreakeven] = useState(price);

	useEffect(() => {
		const pureFee = pureStatusPercent != "" ? 0.01 * parseFloat(pureStatusPercent) : 0;
		const priceAfterPure = price - (price * pureFee)
		console.log(priceAfterPure)
		const maxBuyPrice = priceAfterPure / (1 + (0.01 * feePercent))
		setBreakeven(maxBuyPrice);
	}, [pureStatusPercent, price]);

	return (
		<>
			<h3 className="col-span-2">{paymentName}</h3>
			<p className="border text-center w-full">{`$${(breakeven).toFixed(2)}`}</p>
			<p className="text-center">{}</p>			</>
	)
}
export default PurchaseValueDisplay;
