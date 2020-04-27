<template>
<div>
  <p @click="navigate(showPrev)">{{showPrev}}</p>
  <br>
  <p>{{active}}</p>
  <br>
  <p @click="navigate(showNext)">{{showNext}}</p>
</div>
</template>

<script>
export default {
  props: ['active', 'basketCount'],
  data() {
    return {
    }
  },
  methods: {
    navigate(num) {
      this.$emit('navigate', num);
    }
  },
  computed: {
    showNext() {
      if (this.active >= this.basketCount)
        return (1);
      return this.active + 1;
    },
    showPrev() {
      if (this.active <= 1)
        return (this.basketCount);
      return this.active - 1;
    }
  }
};
</script>

<style>
</style>
