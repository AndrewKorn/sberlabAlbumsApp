<template>
 <div>

   <tr style="display: flex">
     <th scope="row">{{index + 1}}</th>
     <td>{{album.album_name}}</td>
     <td >{{album.songs}}</td>
     <td>
       <button class="btn btn-outline-danger">X</button>
     </td>
   </tr>
<!--   <li>
     <span>
       <strong>{{album.id}}</strong>
       {{album.album_name}}
     </span>
     <button class="btn-outline-danger">
       &times;
    </button>
   </li>-->
 </div>
</template>

<script>
export default {
  name: "TableItem",
  components: {},

  props: {
    album: {
      type: Object,
      required: true
    }
  }
}
</script>

<style>
  li {
    border: 1px solid #ccc;
    display: flex;
    justify-content: space-between;
    padding: .5rem 2rem;
    margin-bottom: 1rem;
  }
</style>