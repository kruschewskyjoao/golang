$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(event) {
  event.preventDefault();

  if($('#senha').val() != $('#confirmar-senha').val()) {
    alert("As senhas são diferentes")
    return
  }
  $.ajax({
    url: "/usuarios",
    method: "POST",
    data: {
      nome: $('#nome').val(),
      email: $('#email').val(),
      nick: $('#nick').val(),
      senha: $('#senha').val()
    }
  }).done(function() {
    alert("usuario cadastrado com sucesso")
  }).fail(function(erro) {
    console.log(erro)
    alert("Erro ao cadastrar")
  });
}