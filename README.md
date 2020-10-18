# Templator

- Generate Html when running locally by making a POST request to `http://localhost:9898/EmailTemplate/` 

- Sampe request payload
```
{
    "product":{
        "name": "Sample Company",
        "link": "https://kwabena.dev",
        "logo": "http://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png",
        "copyright": "Copyright © 2020. All rights reserved.",
        "TroubleText": "Contact us in case of trouble"
    },
    "email": {
        "body": {
        "name": "Samuel",
        "intros": ["With great pleasure do I write to you this.", "Come join our ministry"],
        "actions":[{
            "instructions": "to get started click on the  confirm button to activate your account",
            "invitecode": "99988999",
            "button": {
                "color": "#22BC66",
                "text": "Confirm your account",
				"link": "https://kwabena.dev.com/confirm?token=d9729feb74992cc3482b350163a1a010"
            }
        }],
        "outros": ["Thank you"],
        "greeting": "Hello",
        "signature":"Yours sincerely",
        "title": ""
    }
    }

}
```