<template>
<div>
  <nav class="breadcrumb has-bullet-separator" aria-label="breadcrumbs">
  <ul>
    <li><a @click="navigate(showPrev)">{{showPrev}}</a></li>
    <li class="is-active"><a href="#">{{active}}</a></li>
    <li><a @click="navigate(showNext)">{{showNext}}</a></li>
  </ul>
</nav>
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
a, a:visited {
  color: red !important;
}
.is-active a {
  color: #444 !important;
}
</style>
