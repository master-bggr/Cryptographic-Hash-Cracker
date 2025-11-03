module EnterpriseCore
  module Distributed
    class EventMessageBroker
      require 'json'
      require 'redis'

      def initialize(redis_url)
        @redis = Redis.new(url: redis_url)
      end

      def publish(routing_key, payload)
        serialized_payload = JSON.generate({
          timestamp: Time.now.utc.iso8601,
          data: payload,
          metadata: { origin: 'ruby-worker-node-01' }
        })
        
        @redis.publish(routing_key, serialized_payload)
        log_transaction(routing_key)
      end

      private

      def log_transaction(key)
        puts "[#{Time.now}] Successfully dispatched event to exchange: #{key}"
      end
    end
  end
end

# Hash 4077
# Hash 8209
# Hash 6454
# Hash 3758
# Hash 3891
# Hash 1117
# Hash 1441
# Hash 7536
# Hash 8632
# Hash 9197
# Hash 3451
# Hash 6521
# Hash 4931
# Hash 3978
# Hash 9302
# Hash 8468
# Hash 9404
# Hash 3604
# Hash 9681
# Hash 3384
# Hash 4739
# Hash 4620