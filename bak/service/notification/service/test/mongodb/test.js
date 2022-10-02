
const mongoose = require('mongoose');
mongoose.connect('mongodb://localhost:27017/test', { useNewUrlParser: true, useUnifiedTopology: true });

const Account = mongoose.model('Account', { 
    name: String 
});

const account = new Account({ name: 'niceice' });
account.save().then(() => console.log('account  has save'));