<template>
    <span>
        <v-form v-if='!created' @submit.prevent="signUp">
            <v-text-field label="Username" autofocus outlined v-model="user.username"
                    :color="this.$store.state.accent" background-color="white"
                    counter="25"></v-text-field>
            <v-text-field label="First Name" outlined v-model="user.firstname"
                    :color="this.$store.state.accent" background-color="white"></v-text-field>
            <v-text-field label="Last Name" outlined v-model="user.lastname"
                    :color="this.$store.state.accent" background-color="white"></v-text-field>
            <v-text-field label="Email" outlined type="email" v-model="user.email"
                    :color="this.$store.state.accent" background-color="white"></v-text-field>
            <v-text-field label="Password" outlined type="password" v-model="user.password"
                    :color="this.$store.state.accent" background-color="white"></v-text-field>
            <v-text-field label="Re-enter Password" outlined type="password" v-model="user.secondPassword"
                    :color="this.$store.state.accent" background-color="white"></v-text-field>
            <v-btn type="submit">Sign Up</v-btn>
            <br class="my-5">
        </v-form>
        <h2 v-if='created' class="my-10">Account created successfully!</h2>
        <v-btn small to="/login">Back to login</v-btn>
    </span>
</template>

<script>
export default {
    name: 'SignUpForm',
    data() {
        return {
            created: false,
            user: {
                username: '',
                password: '',
                email: '',
                firstname: '',
                lastname: '',
                secondPassword: ''
            }
        }
    },
    methods: {
        signUp() {
            if(!Object.values(this.user).includes('')) {
                if(this.user.password === this.user.secondPassword){
                    let data = {username: this.user.username,
                        password: this.user.password,
                        email: this.user.email,
                        firstname: this.user.firstname,
                        lastname: this.user.lastname};
                    console.log(data);
                    this.$store.dispatch('CREATE_USER', data)
                    .then(() => {
                        this.created = true;
                        this.$emit('raiseError', '');
                    })
                    .catch(() => {
                        this.created = false;
                        this.$emit('raiseError', 'Failed to create user');
                    })
                }
                else{
                    this.$emit('raiseError', 'The passwords do not match');
                }
            }
            else{
                this.$emit('raiseError', 'No fields can be left empty');
            }
        }
    }
}
</script>