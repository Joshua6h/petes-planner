<template>
    <v-form @submit.prevent="signUp">
        <v-text-field label="Username" autofocus outlined v-model="username"
                :color="this.$store.state.accent" background-color="white"></v-text-field>
        <v-text-field label="Email" outlined type="email" v-model="email"
                :color="this.$store.state.accent" background-color="white"></v-text-field>
        <v-text-field label="Password" outlined type="password" v-model="password"
                :color="this.$store.state.accent" background-color="white"></v-text-field>
        <v-text-field label="Re-enter Password" outlined type="password" v-model="secondPassword"
                :color="this.$store.state.accent" background-color="white"></v-text-field>
        <v-btn type="submit">Sign Up</v-btn>
    </v-form>
</template>

<script>
export default {
    name: 'SignUpForm',
    data() {
        return {
            username: '',
            email: '',
            password: '',
            secondPassword: ''
        }
    },
    methods: {
        signUp() {
            if(this.username && this.email && this.password && this.secondPassword) {
                if(this.password === this.secondPassword){
                    console.log("signed up");
                    console.log(this.username);
                    console.log(this.email);
                    console.log(this.password);
                    console.log(this.secondPassword);
                    this.$emit('raiseError', '');
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