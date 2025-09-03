export function handleTransfer(event: TransferEvent): void {
  let entity = new Transfer(
    event.transaction.hash.concatI32(event.logIndex.toI32())
  )
  entity.transactionHash = event.transaction.hash
  entity.tokenId = event.params.tokenId
  entity.from = event.params.from
  entity.to = event.params.to

  entity.save()
}
