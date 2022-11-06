let calcSection = new Vue({
  el: '#calc_section',
  data: {
    error: "",
    input: {
      x: 0,
      y: 0,
      z: 0
    },
    result: {
      result: 0,
      count: 0
    }
  },
  methods: {
    doCalc(operation) {
      this.error = "";
      this.result.count = 0;
      let req = {
        numbers: [this.input.x, this.input.y, this.input.z]
      }
      axios.post("/api/"+operation, req)
      .then(resp => {
        this.result = resp.data;
      })
      .catch(err => {
        this.error = err;
      });
    },
  }
});