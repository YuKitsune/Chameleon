import React from 'react';
import PropsWithClassName from '../PropsWithClassName';

const PayPalIcon = (props: PropsWithClassName) => {
	return (
		<div className={props.className}>
			<a href="https://www.paypal.com/c2/webapps/mpp/paypal-popup?locale.x=en_C2" title="PayPal" onClick={() => window.open('https://www.paypal.com/c2/webapps/mpp/paypal-popup?locale.x=en_C2','WIPaypal','toolbar=no, location=no, directories=no, status=no, menubar=no, scrollbars=yes, resizable=yes, width=1060, height=700')}><img src="https://www.paypalobjects.com/webstatic/mktg/Logo/pp-logo-100px.png" alt="PayPal Logo"/></a>
		</div>
	);
}

export default PayPalIcon;
