$('#nova-publicacao').on('submit', criarPublicacao)
$(document).on('click', '.curtir-publicacao', curtirPublicacao)
$(document).on('click', '.descurtir-publicacao', descurtirPublicacao)

function criarPublicacao(event) {
  event.preventDefault()

  $.ajax({
    url: "/publicacoes",
    method: "POST",
    data: {
      titulo: $('#titulo').val(),
      conteudo: $('#conteudo').val(),
    }
  }).done(function() {
    window.location = "/home"
  }).fail(function() {
    alert("Erro na criação da publicação")
  })
}

function curtirPublicacao(event) {
  const elementoClicado = $(event.target);
  const publicacaoId = elementoClicado.closest('div').data('publicacao-id')
  elementoClicado.prop('disabled', true)
  $.ajax({
    url: `/publicacoes/${publicacaoId}/curtir`,
    method: "POST",
  }).done(function () {
    const contadorCurtidas = elementoClicado.next('span')
    const quantidadeCurtidas = parseInt(contadorCurtidas.text())
    contadorCurtidas.text(quantidadeCurtidas + 1);

    elementoClicado.addClass('descurtir-publicacao')
    elementoClicado.addClass('text-danger')
    elementoClicado.removeClass('curtir-publicacao')

  }).fail(function() {
    alert("erro ao curtir")
  }).always(function() {
    elementoClicado.prop('disabled', false)
  })

}

function descurtirPublicacao(event) {
  const elementoClicado = $(event.target);
  const publicacaoId = elementoClicado.closest('div').data('publicacao-id')
  elementoClicado.prop('disabled', true)
  $.ajax({
    url: `/publicacoes/${publicacaoId}/descurtir`,
    method: "POST",
  }).done(function () {
    const contadorCurtidas = elementoClicado.next('span')
    const quantidadeCurtidas = parseInt(contadorCurtidas.text())
    contadorCurtidas.text(quantidadeCurtidas - 1);

    elementoClicado.removeClass('descurtir-publicacao')
    elementoClicado.removeClass('text-danger')
    elementoClicado.addClass('curtir-publicacao')

  }).fail(function() {
    alert("erro ao descurtir")
  }).always(function() {
    elementoClicado.prop('disabled', false)
  })

}